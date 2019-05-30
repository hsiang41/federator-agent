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
	Lib "github.com/containers-ai/federatorai-agent/pkg/inputlib"
	Queue "github.com/sheerun/queue"
)

const (
	envVarPrefix = "FEDERATOR_AGENT"
)

var transmitterConfigurationFile string
var agentConfig Agent.Config
var scope *logUtil.Scope
var agentQueue *Queue.Queue

type ScheduleJob struct {
	libPath    string
	configPath string
}

func (s ScheduleJob) Run() {
	p, err := plugin.Open(s.libPath)
	if err != nil {
		scope.Errorf(fmt.Sprintf("Failed to open library: %v", err))
		return
	}

	libName, err := p.Lookup("InputLib")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var libObj Lib.InputLibrary
	libObj, ok := libName.(Lib.InputLibrary)
	if !ok {
		scope.Errorf(fmt.Sprintf("Failed to load library object: %v", err))
		return
	}

	libObj.LoadConfig(s.configPath, scope)
	libObj.Gather()
}

func init() {
	flag.StringVar(&transmitterConfigurationFile, "config", "/etc/alameda/federatorai-agent/transmitter.yml", "File path to transmitter configuration")
	// flag.StringVar(&transmitterConfigurationFile, "config", "/root/goProject/src/github.com/containers-ai/federatorai-agent/etc/transmitter.yml", "File path to transmitter configuration")
	scope = logUtil.RegisterScope("manager", "operator entry point", 0)
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

	for _, v := range agentConfig.InputJobs {
		c.AddFunc(v.ScheduledTaskSpec, func() {
			scope.Debug(fmt.Sprintf("Start cron job: %v", v))
		})
		c.AddJob(v.ScheduledTaskSpec, ScheduleJob{v.LibPath, v.LibConfiguration})
	}

	c.Start()
	defer c.Stop()
	select {}
}
