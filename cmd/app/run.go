package app

import (
	"fmt"
	"errors"
	"encoding/json"
	"flag"
	"os"
	"plugin"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/robfig/cron"
	"github.com/containers-ai/alameda/pkg/utils/log"
	Agent "github.com/containers-ai/federatorai-agent"
	InputLib "github.com/containers-ai/federatorai-agent/pkg/inputlib"
	OutputLib "github.com/containers-ai/federatorai-agent/pkg/outputlib"
	Queue "github.com/sheerun/queue"
	"sync"
)

const (
	envVarPrefix = "FEDERATORAI_AGENT"
)

var (
	agentQueue *Queue.Queue
	cfgPath string
	transmitterConfigurationFile string
	logger = log.RegisterScope("Federatorai-Agent", "Federatorai-Agent", 0)
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "start alameda recommender",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			startAgent()
		},
	}
)

var mutex sync.Mutex

func init() {
	flag.StringVar(&transmitterConfigurationFile, "config", "/etc/alameda/federatorai-agent/transmitter.toml", "File path to transmitter configuration")
	logOpt := log.DefaultOptions()
	logOpt.RotationMaxBackups = 10
	log.Configure(logOpt)
	// flag.StringVar(&transmitterConfigurationFile, "config", "/root/goProject/src/github.com/containers-ai/federatorai-agent/etc/transmitter.toml", "File path to transmitter configuration")
}

func ReadConfig(filename string) (*Agent.Config, error) {
	var agentConfig Agent.Config
	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&agentConfig)
	if err != nil {
		panic(errors.New("Unmarshal configuration failed: " + err.Error()))
	} else {
		logLev, ok := log.StringToLevel(agentConfig.Log.OutputLevel)
		fmt.Println(logger, ok, agentConfig.Log.OutputLevel)
		logger.SetOutputLevel(logLev)
		fmt.Println("Set debug level")
		if transmitterConfBin, err := json.MarshalIndent(agentConfig, "", "  "); err == nil {
			logger.Debug(fmt.Sprintf("Transmitter configuration: %s", string(transmitterConfBin)))
		}
	}
	return &agentConfig, nil
}

func initQueue() {
	agentQueue = Queue.New()
}

type LibType int

const (
	LibTypeInput    LibType = 0
	LibTypeOutput   LibType = 1
)

type ScheduleJob struct {
	LibPath     string
	ConfigPath  string
	LibType     LibType
}

func NewScheduleJob(libPath string, configPath string, libType LibType) *ScheduleJob {
	return &ScheduleJob{LibPath: libPath, ConfigPath: configPath, LibType: libType}
}

func (s *ScheduleJob) Run() {
	var symName string
	mutex.Lock()
	logger.Debugf("Require locker")
	p, err := plugin.Open(s.LibPath)
	if err != nil {
		logger.Errorf(fmt.Sprintf("Failed to open library: %v", err))
		return
	}

	logger.Debugf(fmt.Sprintf("object %p", p))
	if s.LibType == LibTypeInput {
		symName = "InputLib"
	} else
	{
		symName = "OutputLib"
	}
	libName, err := p.Lookup(symName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if s.LibType == LibTypeInput {
		var libObj InputLib.InputLibrary
		libObj, ok := libName.(InputLib.InputLibrary)
		if !ok {
			logger.Errorf(fmt.Sprintf("Failed to load input library object: %v", ok))
			return
		}
		defer libObj.Close()
		//logger.Debugf(fmt.Sprintf("Load config: %s", s.ConfigPath))
		libObj.LoadConfig(&s.ConfigPath, logger)
		libObj.SetAgentQueue(agentQueue)
		libObj.Gather()
	} else if s.LibType == LibTypeOutput {
		var libObj OutputLib.OutputLibrary
		libObj, ok := libName.(OutputLib.OutputLibrary)
		if !ok {
			logger.Errorf(fmt.Sprintf("Failed to load output library object: %v", ok))
			return
		}
		defer libObj.Close()
		//logger.Debugf(fmt.Sprintf("Load config: %s", s.ConfigPath))
		libObj.LoadConfig(s.ConfigPath, logger)
		libObj.SetAgentQueue(agentQueue)
		libObj.Write()

	}
	mutex.Unlock()
	logger.Debugf("Require unlocker")
}

func startAgent() {
	agentConfig, err := ReadConfig(transmitterConfigurationFile)
	if err != nil {
		logger.Fatalf("Failed to read configuration due to %s\n", err)
		return
	}
	initQueue()
	logger.SetLogCallers(agentConfig.Log.SetLogCaller)
	if outputLvl, ok := log.StringToLevel(agentConfig.Log.OutputLevel); ok {
		logger.SetOutputLevel(outputLvl)
	}

	logger.Debug(fmt.Sprintf("input: %v", agentConfig.InputJobs))
	logger.Debug(fmt.Sprintf("output: %v", agentConfig.OutputJobs))

	c := cron.New()

	// For input library
	for _, v := range agentConfig.InputJobs {
		logger.Info(fmt.Sprintf("input: %v", v))
		c.AddFunc(v.ScheduledTaskSpec, func() {
			logger.Debug(fmt.Sprintf("Start cron input job"))
		})
		sJ := NewScheduleJob(v.LibPath, v.LibConfiguration, LibTypeInput)
		logger.Debugf(fmt.Sprintf("function point %p", sJ))
		sJ.Run()
		c.AddJob(v.ScheduledTaskSpec, sJ)
	}

	// For output library
	for _, v := range agentConfig.OutputJobs {
		logger.Info(fmt.Sprintf("output: %v", v))
		c.AddFunc(v.ScheduledTaskSpec, func() {
			logger.Debug(fmt.Sprintf("Start cron output job: %v", v))
		})
		sJ := NewScheduleJob(v.LibPath, v.LibConfiguration, LibTypeInput)
		c.AddJob(v.ScheduledTaskSpec, sJ)
	}

	c.Run()
	c.Start()
	c.Entries()
	defer c.Stop()
	select {}
}