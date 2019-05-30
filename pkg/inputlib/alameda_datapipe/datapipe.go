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
	dataPipeMetrics "github.com/containers-ai/api/datapipe/metrics"
	dataPipeResources "github.com/containers-ai/api/datapipe/resources"

	"github.com/containers-ai/api/datahub/resources"
	"github.com/containers-ai/api/common"
)

type DatapipeConfig struct {
	DataPipe *datahub.Config  `mapstructure:"datapipe"`
}

type inputLib struct {
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

func (d *DataPipeClient) GetNodes() (*dataPipeResources.ListNodesResponse, error) {
	conn, err := grpc.Dial(d.datapipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datapipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datahubClient := dataPipeResources.NewResourcesServiceClient(conn)

	req := dataPipeResources.ListNodesRequest{}

	rep, err := datahubClient.ListNodes(context.Background(), &req)
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to list nodes, %v", err))
		return nil, err
	}

	d.scope.Debug(fmt.Sprintf("get Nodes status: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) GetPods() (*dataPipeResources.ListPodsResponse, error) {
	conn, err := grpc.Dial(d.datapipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datapipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := dataPipeResources.NewResourcesServiceClient(conn)

	req := dataPipeResources.ListPodsRequest{}

	rep, err := datapipeClient.ListPods(context.Background(), &req)
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to list pods, %v", err))
		return nil, err
	}
	d.scope.Debug(fmt.Sprintf("get Pods status: %v", rep.Status.Code))
	return rep, nil
}

func (d *DataPipeClient) GetNodeMetrics(nodeNames []string, timeRange *common.TimeRange) (*dataPipeMetrics.ListNodeMetricsResponse, error) {
	conn, err := grpc.Dial(d.datapipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datapipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
	}
	req := dataPipeMetrics.ListNodeMetricsRequest{
		NodeNames: nodeNames,
		QueryCondition: &qD,
	}

	rep, err := datapipeClient.ListNodeMetrics(context.Background(), &req)
	d.scope.Debug(fmt.Sprintf("get Node metrics status: %v", rep.Status.Code))
	return rep, nil
}

func (d *DataPipeClient) GetPodMetrics(namespaces *resources.NamespacedName, timeRange *common.TimeRange) (*dataPipeMetrics.ListPodMetricsResponse, error) {
	conn, err := grpc.Dial(d.datapipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to connect datahub %s, %v", d.datapipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
	}

	req := dataPipeMetrics.ListPodMetricsRequest{
		NamespacedName: namespaces,
		QueryCondition: &qD,
	}

	rep, err := datapipeClient.ListPodMetrics(context.Background(), &req)
	if err != nil {
		d.scope.Error(fmt.Sprintf("Failed to list pod metrics, %v", err))
		return nil, err
	}
	d.scope.Debug(fmt.Sprintf("get Pod metrics status: %v", rep.Status.Code))
	return rep, nil
}

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.queue = agentQueue
}

func (i inputLib) Gather() error {
	gDClient.GetPods()
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
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

var InputLib inputLib
