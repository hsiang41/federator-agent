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

func init() {
	flag.StringVar(&transmitterConfigurationFile, "config", "/etc/alameda/federatorai-agent/transmitter.toml", "File path to transmitter configuration")
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
	libPath     string
	configPath  string
	libType     LibType
}

func (s ScheduleJob) Run() {
	var symName string
	p, err := plugin.Open(s.libPath)
	if err != nil {
		logger.Errorf(fmt.Sprintf("Failed to open library: %v", err))
		return
	}

	if s.libType == LibTypeInput {
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

	if s.libType == LibTypeInput {
		var libObj InputLib.InputLibrary
		libObj, ok := libName.(InputLib.InputLibrary)
		if !ok {
			logger.Errorf(fmt.Sprintf("Failed to load input library object: %v", err))
			return
		}

		libObj.LoadConfig(s.configPath, logger)
		libObj.SetAgentQueue(agentQueue)
		libObj.Gather()
	} else if s.libType == LibTypeOutput {
		var libObj OutputLib.OutputLibrary
		libObj, ok := libName.(OutputLib.OutputLibrary)
		if !ok {
			logger.Errorf(fmt.Sprintf("Failed to load output library object: %v", err))
			return
		}

		libObj.LoadConfig(s.configPath, logger)
		libObj.SetAgentQueue(agentQueue)
		libObj.Write()
	}
}

func startAgent() {
	agentConfig, err := ReadConfig(transmitterConfigurationFile)
	if err != nil {
		logger.Fatalf("Failed to read configuration due to %s\n", err)
		return
	}
	initQueue()
	logger.SetLogCallers(agentConfig.Log.SetLogCallers)
	if outputLvl, ok := log.StringToLevel(agentConfig.Log.OutputLevel); ok {
		logger.SetOutputLevel(outputLvl)
	}
	if stacktraceLevel, ok := log.StringToLevel(agentConfig.Log.StackTraceLevel); ok {
		logger.SetStackTraceLevel(stacktraceLevel)
	}

	logger.Debug(fmt.Sprintf("input: %v", agentConfig.InputJobs))
	logger.Debug(fmt.Sprintf("output: %v", agentConfig.OutputJobs))

	c := cron.New()

	// For input library
	for _, v := range agentConfig.InputJobs {
		c.AddFunc(v.ScheduledTaskSpec, func() {
			logger.Debug(fmt.Sprintf("Start cron input job: %v", v))
		})
		c.AddJob(v.ScheduledTaskSpec, ScheduleJob{v.LibPath, v.LibConfiguration, LibTypeInput})
	}

	// For output library
	for _, v := range agentConfig.OutputJobs {
		c.AddFunc(v.ScheduledTaskSpec, func() {
			logger.Debug(fmt.Sprintf("Start cron output job: %v", v))
		})
		c.AddJob(v.ScheduledTaskSpec, ScheduleJob{v.LibPath, v.LibConfiguration, LibTypeOutput})
	}

	c.Start()
	defer c.Stop()
	select {}
}