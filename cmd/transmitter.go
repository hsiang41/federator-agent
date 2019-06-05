package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"plugin"
	"strings"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	Agent "github.com/containers-ai/federatorai-agent"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	InputLib "github.com/containers-ai/federatorai-agent/pkg/inputlib"
	OutputLib "github.com/containers-ai/federatorai-agent/pkg/outputlib"
	Queue "github.com/sheerun/queue"
)

const (
	envVarPrefix = "FEDERATOR_AGENT"
)

var transmitterConfigurationFile string
var agentConfig Agent.Config
var scope *logUtil.Scope
var agentQueue *Queue.Queue

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
		scope.Errorf(fmt.Sprintf("Failed to open library: %v", err))
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
			scope.Errorf(fmt.Sprintf("Failed to load input library object: %v", err))
			return
		}

		libObj.LoadConfig(s.configPath, scope)
		libObj.SetAgentQueue(agentQueue)
		libObj.Gather()
	} else if s.libType == LibTypeOutput {
		var libObj OutputLib.OutputLibrary
		libObj, ok := libName.(OutputLib.OutputLibrary)
		if !ok {
			scope.Errorf(fmt.Sprintf("Failed to load output library object: %v", err))
			return
		}

		libObj.LoadConfig(s.configPath, scope)
		libObj.SetAgentQueue(agentQueue)
		libObj.Write()
	}
}

func init() {
	flag.StringVar(&transmitterConfigurationFile, "config", "/etc/alameda/federatorai-agent/transmitter.toml", "File path to transmitter configuration")
	// flag.StringVar(&transmitterConfigurationFile, "config", "/root/goProject/src/github.com/containers-ai/federatorai-agent/etc/transmitter.toml", "File path to transmitter configuration")
	scope = logUtil.RegisterScope("federatorai-agent", "operator entry point", 0)
}

func initConfiguration() {
	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(transmitterConfigurationFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&agentConfig)
	if err != nil {
		panic(errors.New("Unmarshal configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(agentConfig, "", "  "); err == nil {
			scope.Debug(fmt.Sprintf("Transmitter configuration: %s", string(transmitterConfBin)))
		}
	}
}

func initQueue() {
	agentQueue = Queue.New()
}

func main() {
	initConfiguration()
	initQueue()
	scope.SetLogCallers(agentConfig.Log.SetLogCallers)
	if outputLvl, ok := logUtil.StringToLevel(agentConfig.Log.OutputLevel); ok {
		scope.SetOutputLevel(outputLvl)
	}
	if stacktraceLevel, ok := logUtil.StringToLevel(agentConfig.Log.StackTraceLevel); ok {
		scope.SetStackTraceLevel(stacktraceLevel)
	}

	scope.Debug(fmt.Sprintf("input: %v", agentConfig.InputJobs))
	scope.Debug(fmt.Sprintf("output: %v", agentConfig.OutputJobs))

	c := cron.New()

	// For input library
	for _, v := range agentConfig.InputJobs {
		c.AddFunc(v.ScheduledTaskSpec, func() {
			scope.Debug(fmt.Sprintf("Start cron input job: %v", v))
		})
		c.AddJob(v.ScheduledTaskSpec, ScheduleJob{v.LibPath, v.LibConfiguration, LibTypeInput})
	}

	// For output library
	for _, v := range agentConfig.OutputJobs {
		c.AddFunc(v.ScheduledTaskSpec, func() {
			scope.Debug(fmt.Sprintf("Start cron output job: %v", v))
		})
		c.AddJob(v.ScheduledTaskSpec, ScheduleJob{v.LibPath, v.LibConfiguration, LibTypeOutput})
	}

	c.Start()
	defer c.Stop()
	select {}
}
