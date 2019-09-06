package common

import (
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	DataHubCommon "github.com/containers-ai/api/common"
)

type QueueType int

const (
	QueueTypePod            QueueType = 0
	QueueTypeNode           QueueType = 1
	QueueTypePodMetrics     QueueType = 2
	QueueTypeNodeMetrics    QueueType = 3
)

type AgentQueueItem struct {
	QueueType QueueType
	DataItem  interface{}
}

type InfluxField struct {
	Name string
	Type DataHubCommon.DataType
	Default interface{}
}

type DataSourceConverter interface {
	GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error)
}

type TargetData interface {
	IsTagKey(string) bool
	GetFieldKey(string) *InfluxField
	GetSourceData() interface{}
}

type ConvertInt     int

const (
	ConvertDatahub  = iota
	ConvertInflux
	ConvertPrometheus
)