package main

import (
	"errors"
	"strings"
	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
	"fmt"
)

type inputLib struct {
}

var gDClient *datapipe.DataPipeClient

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
}

func (i inputLib) Gather() error {
	/*
	_, err := gDClient.GetPods()
	if err != nil {
		gDClient.Scope.Error(fmt.Sprintf("Failed to get pods info, %v", err))
	}
	*/
	/*
	qItem := fmt.Sprintf("%d gather", time.Now().Unix())

	agentQ := &AgentCommon.AgentQueueItem{AgentCommon.QueueTypePod, qItem}
	if gDClient.Queue != nil {
		gDClient.Queue.Append(agentQ)
	}
	*/
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = datapipe.NewDataPipeClient()
	gDClient.Scope = scope

	viper.SetEnvPrefix("InputLibDataHub")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library datahub configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&gDClient.DataPipe)
	if err != nil {
		panic(errors.New("Unmarshal input library datahub configuration failed: " + err.Error()))
	} else {
		/*
		if transmitterConfBin, err := json.MarshalIndent(gDClient.datapipe, "", "  "); err == nil {
			gDClient.scope.Infof(fmt.Sprintf("Input library datahub configuration: %s", string(transmitterConfBin)))
		}
		*/
	}

	return nil
}

var InputLib inputLib
