package datapipe

import (
	"context"
	"github.com/sheerun/queue"
	"google.golang.org/grpc"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/alameda/operator/datahub"
	dataPipeMetrics "github.com/containers-ai/api/datapipe/metrics"
	dataPipeResources "github.com/containers-ai/api/datapipe/resources"
	"github.com/containers-ai/api/datapipe/predictions"
	"github.com/containers-ai/api/datahub/resources"
	"github.com/containers-ai/api/datahub/metrics"
	"github.com/containers-ai/api/common"
	"fmt"
	datahub_v1a1pha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"unsafe"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type DatapipeConfig struct {
	DataPipe            *datahub.Config  `mapstructure:"datapipe"`
	DataGranularitySec	int32            `mapstructure:"data_granularity_sec"`
	DataAmountInitSec   int32            `mapstructure:"data_amount_init_sec"`
	DataAmountSec       int32            `mapstructure:"data_amount_sec"`
}

type DataPipeClient struct {
	DataPipe DatapipeConfig
	Scope    *logUtil.Scope
	Queue    *queue.Queue
}

func NewDataPipeClient() *DataPipeClient {
	return &DataPipeClient{}
}

func (d *DataPipeClient) GetNodes() (*datahub_v1a1pha1.ListNodesResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := datahub_v1a1pha1.NewDatahubServiceClient(conn)

	req := datahub_v1a1pha1.ListAlamedaNodesRequest{}

	rep, err := datapipeClient.ListAlamedaNodes(context.Background(), &req)
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to list nodes, %v", err))
		return nil, err
	}

	d.Scope.Debug(fmt.Sprintf("get Nodes: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) GetPods() (*datahub_v1a1pha1.ListPodsResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := datahub_v1a1pha1.NewDatahubServiceClient(conn)

	req := datahub_v1a1pha1.ListAlamedaPodsRequest{}

	rep, err := datapipeClient.ListAlamedaPods(context.Background(), &req)
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to list pods, %v", err))
		return nil, err
	}
	d.Scope.Debug(fmt.Sprintf("get Pods: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) GetNodeMetrics(nodeNames []string, timeRange *common.TimeRange, limit uint64) (*dataPipeMetrics.ListNodeMetricsResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
		Order: common.QueryCondition_DESC,
	}
	if limit > 0 {
		qD.Limit = limit
	}
	req := dataPipeMetrics.ListNodeMetricsRequest{
		NodeNames: nodeNames,
		QueryCondition: &qD,
	}

	rep, err := datapipeClient.ListNodeMetrics(context.Background(), &req)
	if rep.Status.Code != 0 {
		d.Scope.Errorf(fmt.Sprintf("Failed to list %v node metrics(%d): %s", nodeNames, rep.Status.Code, rep.Status.Message))
		return rep, err
	}
	d.Scope.Debugf(fmt.Sprintf("Succeed to list node metrics: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) GetPodMetrics(namespaces *resources.NamespacedName, timeRange *common.TimeRange, limit uint64) (*dataPipeMetrics.ListPodMetricsResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()
	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
		Order: common.QueryCondition_DESC,
	}
	if limit > 0 {
		qD.Limit = limit
	}

	req := dataPipeMetrics.ListPodMetricsRequest{
		NamespacedName: namespaces,
		QueryCondition: &qD,
	}

	rep, err := datapipeClient.ListPodMetrics(context.Background(), &req)
	if err != nil {
		d.Scope.Errorf(fmt.Sprintf("Failed to list %s pods metrics, %v", namespaces.Namespace, err))
		return nil, err
	}
	if rep.Status.Code != 0 {
		d.Scope.Errorf(fmt.Sprintf("Failed to list %s pods metrics(%d): %s", namespaces.Namespace, rep.Status.Code, rep.Status.Message))
		return rep, err
	}
	d.Scope.Debugf(fmt.Sprintf("Succeed to list pods metrics: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) CreatePods(pods []*resources.Pod) error {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return err
	}
	defer conn.Close()

	datapipeClient := dataPipeResources.NewResourcesServiceClient(conn)
	req := &dataPipeResources.CreatePodsRequest{
		Pods:pods,
	}
	// TODO: need parese status result
	status, err := datapipeClient.CreatePods(context.Background(), req)
	if err != nil {
		return err
	}
	d.Scope.Debug(fmt.Sprintf("Create pods status: %v", status))
	return nil
}

func (d *DataPipeClient) CreateNodes(nodes []*resources.Node) error {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return err
	}
	defer conn.Close()

	datapipeClient := dataPipeResources.NewResourcesServiceClient(conn)
	req := &dataPipeResources.CreateNodesRequest{
		Nodes: nodes,
	}
	// TODO: need parese status result
	status, err := datapipeClient.CreateNodes(context.Background(), req)
	if err != nil {
		return err
	}
	d.Scope.Debug(fmt.Sprintf("Create nodes status: %v", status))
	return nil
}

func (d *DataPipeClient) CreateNodeMetrics(nodesMetrics []*metrics.NodeMetric) error {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return err
	}
	defer conn.Close()

	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	req := &dataPipeMetrics.CreateNodeMetricsRequest{
		NodeMetrics: nodesMetrics,
	}
	// TODO: need parese status result
	status, err := datapipeClient.CreateNodeMetrics(context.Background(), req)
	if err != nil {
		return err
	}
	if status.Code != 0 {
		d.Scope.Errorf(fmt.Sprintf("Failed to create nodes metrics(%d): %s", status.Code, status.Message))
		return err
	}
	d.Scope.Debug(fmt.Sprintf("Create nodes metrics status: %v", status))
	return nil
}

func (d *DataPipeClient) CreatePodMetrics(podsMetrics []*metrics.PodMetric) error {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return err
	}
	defer conn.Close()

	datapipeClient := dataPipeMetrics.NewMetricsServiceClient(conn)
	req := &dataPipeMetrics.CreatePodMetricsRequest{
		PodMetrics: podsMetrics,
	}
	// TODO: need parese status result
	status, err := datapipeClient.CreatePodMetrics(context.Background(), req)
	if err != nil {
		return err
	}
	if status.Code != 0 {
		d.Scope.Errorf(fmt.Sprintf("Failed to create pods metrics(%d): %s", status.Code, status.Message))
		return err
	}
	d.Scope.Debug(fmt.Sprintf("Create pods metrics status: %v", status))
	return nil
}

func (d *DataPipeClient) GetNodesPredictions(nodesName []string, timeRange *common.TimeRange, limit uint64) (*predictions.ListNodePredictionsResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()

	datapipeClient := predictions.NewPredictionsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
		Order: common.QueryCondition_DESC,
	}
	if limit > 0 {
		qD.Limit = limit
	}
	req := &predictions.ListNodePredictionsRequest{
		NodeNames: nodesName,
		QueryCondition: &qD,
	}
	rep, err := datapipeClient.ListNodePredictions(context.Background(), req)
	if err != nil {
		return nil, err
	}
	d.Scope.Debug(fmt.Sprintf("List nodes predictions status: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) GetPodsPredictions(namespaces *resources.NamespacedName, timeRange *common.TimeRange, limit uint64) (*predictions.ListPodPredictionsResponse, error) {
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()

	datapipeClient := predictions.NewPredictionsServiceClient(conn)
	qD := common.QueryCondition{
		TimeRange: timeRange,
		Order: common.QueryCondition_DESC,
	}
	if limit > 0 {
		qD.Limit = limit
	}
	req := &predictions.ListPodPredictionsRequest{
		NamespacedName: namespaces,
		QueryCondition: &qD,
	}
	rep, err := datapipeClient.ListPodPredictions(context.Background(), req)
	if err != nil {
		return nil, err
	}
	d.Scope.Debug(fmt.Sprintf("List pods predictions status: %v", rep))
	return rep, nil
}

func (d *DataPipeClient) ListPodRecommendations(namespaces *resources.NamespacedName, timeRange *common.TimeRange, limit uint64) (*datahub_v1a1pha1.ListPodRecommendationsResponse, error) {
	var qD datahub_v1a1pha1.QueryCondition
	conn, err := grpc.Dial(d.DataPipe.DataPipe.Address, grpc.WithInsecure())
	if err != nil {
		d.Scope.Error(fmt.Sprintf("Failed to connect datapipe %s, %v", d.DataPipe.DataPipe.Address, err))
		return nil, err
	}
	defer conn.Close()

	datapipeClient := datahub_v1a1pha1.NewDatahubServiceClient(conn)
	qD.Order = datahub_v1a1pha1.QueryCondition_DESC
	if timeRange != nil {
		qD.TimeRange = (*datahub_v1a1pha1.TimeRange)(unsafe.Pointer(timeRange))
	}
	if limit > 0 {
		qD.Limit = limit
	}
	d.Scope.Debugf(fmt.Sprintf("Query condition: %v", qD))
	req := &datahub_v1a1pha1.ListPodRecommendationsRequest{
		NamespacedName: (*datahub_v1a1pha1.NamespacedName)(unsafe.Pointer(namespaces)),
		QueryCondition: &qD,
	}
	rep, err := datapipeClient.ListPodRecommendations(context.Background(), req)
	if err != nil {
		d.Scope.Errorf(fmt.Sprintf("Failed to list %s pods recommendation: %v", ((*datahub_v1a1pha1.NamespacedName)(unsafe.Pointer(namespaces))).Namespace, err))
		return nil, err
	}
	if rep.Status.Code != 0 {
		d.Scope.Errorf(fmt.Sprintf("Failed to list %s pods recommendation(%d): %v",
			((*datahub_v1a1pha1.NamespacedName)(unsafe.Pointer(namespaces))).Namespace,
			rep.Status.Code, rep.Status.GetMessage()))
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to list %s pods recommendation(%d): %v",
			((*datahub_v1a1pha1.NamespacedName)(unsafe.Pointer(namespaces))).Namespace,
			rep.Status.Code, rep.Status.GetMessage()))
	}
	d.Scope.Debugf(fmt.Sprintf("List %s pods recommendations status: %v", ((*datahub_v1a1pha1.NamespacedName)(unsafe.Pointer(namespaces))).Namespace, rep))
	return rep, nil
}

func (d *DataPipeClient) ConvertPodNamespace(pod *datahub_v1a1pha1.Pod) (*resources.NamespacedName) {
	if pod == nil {
		return &resources.NamespacedName{}
	}
	return &resources.NamespacedName{
		pod.NamespacedName.Namespace,
		pod.NamespacedName.Name,
		pod.NamespacedName.XXX_NoUnkeyedLiteral,
		pod.NamespacedName.XXX_unrecognized,
		pod.NamespacedName.XXX_sizecache,
	}
}