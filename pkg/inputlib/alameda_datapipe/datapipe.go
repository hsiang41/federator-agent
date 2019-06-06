package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"strings"
	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
)

type inputLib struct {
}

var gDClient *datapipe.DataPipeClient

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
}

func (i inputLib) Gather() error {

	// For Pods metrics
	lsPodResp, err := gDClient.GetPods()
	if err != nil {
		gDClient.Scope.Error(fmt.Sprintf("Failed to get pods info, %v", err))
	}

	if lsPodResp != nil {
		for _, v := range lsPodResp.Pods {
			tr := utils.GetTimeRange(nil, nil, gDClient.DataPipe.DataAmountSec, true, gDClient.DataPipe.DataGranularitySec)
			namespace := gDClient.ConvertPodNamespace(v)
			podMetrics, err := gDClient.GetPodMetrics(namespace, tr)
			if err != nil {
				gDClient.Scope.Error(fmt.Sprintf("Failed to get pod metrics, %v", err))
				continue
			}
			podM := podMetrics.GetPodMetrics()
			if podM != nil {
				err := gDClient.CreatePodMetrics(podM)
				if err != nil {
					gDClient.Scope.Errorf(fmt.Sprintf("Failed to create pods metrics, %v", err))
				} else {
					gDClient.Scope.Debugf(fmt.Sprintf("Succeed to create pod metrics"))
				}

			}
		}
	}

	// For nodes metrics
	lsNodeResp, err := gDClient.GetNodes()
	if err != nil {
		gDClient.Scope.Error(fmt.Sprintf("Failed to get nodes info, %v", err))
	}

	if lsNodeResp != nil {
		var nodesName [] string
		for _, v := range lsNodeResp.Nodes {
			nodesName = append(nodesName, v.GetName())
		}
		if len(nodesName) > 0 {
			tr := utils.GetTimeRange(nil, nil, gDClient.DataPipe.DataAmountSec, true, gDClient.DataPipe.DataGranularitySec)
			nodeMetrics, err := gDClient.GetNodeMetrics(nodesName, tr)
			if err != nil {
				gDClient.Scope.Error(fmt.Sprintf("Failed to get nodes metrics, %v", err))
			}
			if nodeMetrics != nil {
				err := gDClient.CreateNodeMetrics(nodeMetrics.GetNodeMetrics())
				if err != nil {
					gDClient.Scope.Errorf(fmt.Sprintf("Failed to create nodes metrics, %v", err))
				} else {
					gDClient.Scope.Debugf(fmt.Sprintf("Succeed to create nodes metrics"))
				}
			}
		}
	}
	
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
		if transmitterConfBin, err := json.MarshalIndent(gDClient.DataPipe, "", "  "); err == nil {
			gDClient.Scope.Debugf(fmt.Sprintf("Input library datahub configuration: %s", string(transmitterConfBin)))
		}
	}
	return nil
}

var InputLib inputLib
