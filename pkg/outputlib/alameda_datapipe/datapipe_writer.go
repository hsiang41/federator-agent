package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"encoding/json"
	"github.com/sheerun/queue"
	AgentQueue "github.com/containers-ai/federatorai-agent/pkg"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
)

type outputlib struct {
}

var gDClient *datapipe.DataPipeClient

func (i outputlib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
}

func (i outputlib) Write() error {
	if gDClient.Queue == nil {
		return nil
	}

	if gDClient.Queue.Length() == 0 {
		return nil
	}

	q := gDClient.Queue.Pop()
	if q != nil {
		gDClient.Scope.Debug(fmt.Sprintf("Output lib write %s", q.(*AgentQueue.AgentQueueItem).DataItem.(string)))
		q = nil
	}
	return nil
}

func (i outputlib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = datapipe.NewDataPipeClient()
	gDClient.Scope = scope

	viper.SetEnvPrefix("OutputLibDataHub")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read output library datahub configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&gDClient.DataPipe)
	if err != nil {
		panic(errors.New("Unmarshal output library datahub configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(gDClient.DataPipe, "", "  "); err == nil {
			gDClient.Scope.Infof(fmt.Sprintf("Output library datahub configuration: %s", string(transmitterConfBin)))
		}
	}

	return nil
}

var OutputLib outputlib
