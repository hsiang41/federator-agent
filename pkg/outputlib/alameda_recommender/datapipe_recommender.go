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
	rcWriter "github.com/containers-ai/federatorai-agent/pkg/crwriter"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type outputlib struct {
}

var gDClient *datapipe.DataPipeClient

func (i outputlib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
}

func (i outputlib) Write() error {
	gDClient.Scope.Debugf(fmt.Sprintf("output library - Write"))
	crWriter, err := rcWriter.NewCrWriter(gDClient.Scope)
	if err != nil {
		gDClient.Scope.Errorf(fmt.Sprintf("Failed to new cr writer object: %v", err))
	}
	// List pod recommendation
	lsPodResp, err := gDClient.GetPods()
	if err != nil {
		gDClient.Scope.Error(fmt.Sprintf("Failed to get pods list, %v", err))
		return err
	}

	if lsPodResp.Status.Code != 0 {
		gDClient.Scope.Errorf(fmt.Sprintf("Failed to get pods list(%d), %v", lsPodResp.Status.Code, lsPodResp.Status.GetMessage()))
		return status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get pods list(%d), %v", lsPodResp.Status.Code, lsPodResp.Status.GetMessage()))
	}

	if lsPodResp != nil {
		for _, v := range lsPodResp.Pods {
			// tr := utils.GetTimeRange(nil, nil, gDClient.DataPipe.DataAmountSec, true, gDClient.DataPipe.DataGranularitySec)
			namespace := gDClient.ConvertPodNamespace(v)
			podRecommendations, err := gDClient.ListPodRecommendations(namespace, nil, 1)
			if err != nil {
				gDClient.Scope.Errorf(fmt.Sprintf("Failed to get %s pod recommendations, %v", namespace, err))
				continue
			}
			if podRecommendations == nil {
				gDClient.Scope.Warnf(fmt.Sprintf("Current %s pod did not have recommendations", namespace))
			}
			if podRecommendations.Status.Code != 0 {
				gDClient.Scope.Errorf(fmt.Sprintf("Failed to get %s pod recommendations(%d), %s", namespace, podRecommendations.Status.Code, podRecommendations.Status.GetMessage()))
			} else {
				podRC := podRecommendations.PodRecommendations
				if podRC != nil {
					gDClient.Scope.Debugf(fmt.Sprintf("Start to write %s pod recommendations: %v", namespace, podRC))
					crWriter.CreatePodRecommendations(context.Background(), podRC)
				} else {
					gDClient.Scope.Warnf(fmt.Sprintf("Current %s pod did not have CR", namespace))
				}
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
