package datahub

import (
	"context"
	"google.golang.org/grpc"
	"fmt"
	"encoding/json"
	"github.com/sheerun/queue"
	"github.com/containers-ai/alameda/operator/datahub"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	datahubV1a1pha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/golang/protobuf/ptypes"
)

type DataHubConfig struct {
	DataHub            *datahub.Config   `mapstructure:"datapipe"`
	DataGranularitySec	int32            `mapstructure:"data_granularity_sec"`
	DataAmountInitSec   int32            `mapstructure:"data_amount_init_sec"`
	DataAmountSec       int32            `mapstructure:"data_amount_sec"`
}

type DataHubClient struct {
	DataHub DataHubConfig
	Scope    *logUtil.Scope
	Queue    *queue.Queue
}

func NewDataHubClient() *DataHubClient {
	return &DataHubClient{}
}

func (d *DataHubClient) GetNodes() (*datahubV1a1pha1.ListNodesResponse, error) {
	conn, err := grpc.Dial(d.DataHub.DataHub.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataHub.DataHub.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := datahubV1a1pha1.NewDatahubServiceClient(conn)

	req := datahubV1a1pha1.ListAlamedaNodesRequest{}

	rep, err := datapipeClient.ListAlamedaNodes(context.Background(), &req)
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to list nodes, %v", err))
		return nil, err
	}

	if jsonNodes, err := json.MarshalIndent(rep, "", "  "); err == nil {
		d.Scope.Debugf(fmt.Sprintf("get Nodes: %s", string(jsonNodes)))
	}
	return rep, nil
}

func (d *DataHubClient) WriteRawData(rawData *datahubV1a1pha1.WriteRawdataRequest) error {
	conn, err := grpc.Dial(d.DataHub.DataHub.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataHub.DataHub.Address, err))
		return err
	}
	defer conn.Close()
	datapipeClient := datahubV1a1pha1.NewDatahubServiceClient(conn)

	rep, err := datapipeClient.WriteRawdata(context.Background(), rawData)
	if err != nil {
		return err
	}
	if rep.Code != 0 {
		return status.Error(codes.Internal, rep.Message)
	}
	return nil
}

func (d *DataHubClient) SetDataPipeAddress(address string) {
	if d.DataHub.DataHub == nil {
		d.DataHub.DataHub = new(datahub.Config)
	}
	d.DataHub.DataHub.Address = address
}

func (d *DataHubClient) SendNotifyEvent(id string, clusterid string, source *datahubV1a1pha1.EventSource, eventType datahubV1a1pha1.EventType, eventLevel datahubV1a1pha1.EventLevel, k8sObject *datahubV1a1pha1.K8SObjectReference, message string, data string) error {
	var dpEvents []*datahubV1a1pha1.Event
	dpEvent := &datahubV1a1pha1.Event{
		Time: ptypes.TimestampNow(),
		Id: id,
		ClusterId: clusterid,
		Source: source,
		Type: eventType,
		Version: datahubV1a1pha1.EventVersion_EVENT_VERSION_V1,
		Level: eventLevel,
		Subject: k8sObject,
		Message: message,
		Data: data,
	}
	dpEvents = append(dpEvents, dpEvent)

	conn, err := grpc.Dial(d.DataHub.DataHub.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataHub.DataHub.Address, err))
		return err
	}
	defer conn.Close()
	datapipeClient := datahubV1a1pha1.NewDatahubServiceClient(conn)

	req := &datahubV1a1pha1.CreateEventsRequest{Events: dpEvents}

	rep, err := datapipeClient.CreateEvents(context.Background(), req)
	if err != nil {
		return err
	}
	if rep.Code != 0 {
		return status.Error(codes.Internal, rep.Message)
	}
	return nil
}