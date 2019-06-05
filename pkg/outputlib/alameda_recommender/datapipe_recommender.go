package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"encoding/json"
	"github.com/sheerun/queue"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	rcWriter "github.com/containers-ai/federatorai-agent/pkg/crwriter"
	"context"
)

type outputlib struct {
}

var gDClient *datapipe.DataPipeClient

func (i outputlib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
}

func (i outputlib) Write() error {
	crWriter, err := rcWriter.NewCrWriter(gDClient.Scope)
	if err != nil {
		gDClient.Scope.Errorf(fmt.Sprintf("Failed to new cr writer object: %v", err))
	}
	// List pod recommendation
	lsPodResp, err := gDClient.GetPods()
	if err != nil {
		gDClient.Scope.Error(fmt.Sprintf("Failed to get pods info, %v", err))
	}

	if lsPodResp != nil {
		for _, v := range lsPodResp.Pods {
			tr := utils.GetTimeRange(nil, nil, gDClient.DataPipe.DataAmountSec, true, gDClient.DataPipe.DataGranularitySec)
			namespace := gDClient.ConvertPodNamespace(v)
			podRecommendations, err := gDClient.ListPodRecommendations(namespace, tr)
			if err != nil {
				gDClient.Scope.Error(fmt.Sprintf("Failed to get pod metrics, %v", err))
				continue
			}
			podRC := podRecommendations.PodRecommendations
			if podRC != nil {
				crWriter.CreatePodRecommendations(context.Background(), podRC)
			}
		}
	}

	return nil
}

func (i outputlib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = datapipe.NewDataPipeClient()
	gDClient.Scope = scope

	viper.SetEnvPrefix("recommender")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read output library recommender configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&gDClient.DataPipe)
	if err != nil {
		panic(errors.New("Unmarshal output library recommender configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(gDClient.DataPipe, "", "  "); err == nil {
			gDClient.Scope.Infof(fmt.Sprintf("Output library recommender configuration: %s", string(transmitterConfBin)))
		}
	}

	return nil
}

var OutputLib outputlib
