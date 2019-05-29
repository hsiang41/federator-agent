package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"encoding/json"
	"github.com/containers-ai/alameda/operator/datahub"
	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
)

type DatahubConfig struct {
	DataHub *datahub.Config  `mapstructure:"datahub"`
}

type inputLib struct {
}

type DataHubClient struct {
	datahub     DatahubConfig
	scope       *logUtil.Scope
	queue       *queue.Queue
}

var gDClient *DataHubClient

func NewDataHubClient() *DataHubClient {
	return &DataHubClient{}
}

func (d *DataHubClient) GetNodes() {
	conn, err := grpc.Dial(d.datahub.DataHub.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datahub.DataHub.Address, err))
		return
	}
	defer conn.Close()
	datahubClient := datahub_v1alpha1.NewDatahubServiceClient(conn)

	req := datahub_v1alpha1.ListNodesRequest{}

	rep, err := datahubClient.ListNodes(context.Background(), &req)
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to list nodes, %v", err))
		return
	}

	d.scope.Info(fmt.Sprintf("get Nodes status: %v", rep))
}

func (d *DataHubClient) GetPods() {
	conn, err := grpc.Dial(d.datahub.DataHub.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datahub.DataHub.Address, err))
		return
	}
	defer conn.Close()
	datahubClient := datahub_v1alpha1.NewDatahubServiceClient(conn)

	req := datahub_v1alpha1.ListAlamedaPodsRequest{}

	rep, err := datahubClient.ListAlamedaPods(context.Background(), &req)
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to list pods, %v", err))
		return
	}

	d.scope.Info(fmt.Sprintf("get Pods status: %v", rep.Status.Code))
}


func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.queue = agentQueue
}

func (i inputLib) Gather() error {
	gDClient.GetPods()
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = NewDataHubClient()
	gDClient.scope = scope

	viper.SetEnvPrefix("InputLibDataHub")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library datahub configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&gDClient.datahub)
	if err != nil {
		panic(errors.New("Unmarshal input library datahub configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(gDClient.datahub, "", "  "); err == nil {
			gDClient.scope.Infof(fmt.Sprintf("Input library datahub configuration: %s", string(transmitterConfBin)))
		}
	}

	return nil
}

var InputLib inputLib
