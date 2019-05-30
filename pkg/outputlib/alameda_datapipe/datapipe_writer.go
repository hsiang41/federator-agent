package main

import (
	"errors"
	"github.com/sheerun/queue"
	"github.com/containers-ai/alameda/operator/datahub"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"encoding/json"
)

type DatapipeConfig struct {
	DataPipe *datahub.Config  `mapstructure:"datapipe"`
}

type outputlib struct {
}

type DataPipeClient struct {
	datapipe    DatapipeConfig
	scope       *logUtil.Scope
	queue       *queue.Queue
}

var gDClient *DataPipeClient

func NewDataPipeClient() *DataPipeClient {
	return &DataPipeClient{}
}

func (i outputlib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.queue = agentQueue
}

func (i outputlib) Write() error {
	return nil
}

func (i outputlib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = NewDataPipeClient()
	gDClient.scope = scope

	viper.SetEnvPrefix("InputLibDataHub")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library datahub configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&gDClient.datapipe)
	if err != nil {
		panic(errors.New("Unmarshal input library datahub configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(gDClient.datapipe, "", "  "); err == nil {
			gDClient.scope.Infof(fmt.Sprintf("Input library datahub configuration: %s", string(transmitterConfBin)))
		}
	}

	return nil
}

var OutputLib outputlib
