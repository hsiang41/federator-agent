package InfluxConvert

import (
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/containers-ai/federatorai-agent/pkg"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert/datahub"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert/influx"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert/prometheus"
)


type Influx struct {
	DatabaseName    string
	MeasurementName string
	TagKeys         []string
	FieldKeys       []*common.InfluxField
	SourceData      interface{}
	Convert         *common.DataSourceConverter
	ConvertType     common.ConvertInt
}

func NewInflux(databaseName string, measurementName string, tags []string, fields []*common.InfluxField, sourceData interface{}, convert common.ConvertInt) *Influx {
	return &Influx{
		DatabaseName:    databaseName,
		MeasurementName: measurementName,
		TagKeys:         tags, FieldKeys: fields, SourceData: sourceData, ConvertType: convert}
}

func (n *Influx)IsTagKey(keyName string) bool {
	if n.TagKeys == nil {
		return false
	}
	for _, n := range n.TagKeys {
		if n == keyName {
			return true
		}
	}
	return false
}

func (n *Influx)GetFieldKey(keyName string) *common.InfluxField {
	for _, n := range n.FieldKeys {
		if n.Name == keyName {
			return n
		}
	}
	return nil
}

func (n *Influx) GetSourceData () interface{} {
	return n.SourceData
}

func (n *Influx)GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error) {
	var dbConvert common.DataSourceConverter
	switch n.ConvertType {
	case common.ConvertDatahub:
		dbConvert = datahub.NewDatahubConverter(n, n.DatabaseName, n.MeasurementName)
	case common.ConvertInflux:
		dbConvert = influx.NewInfluxConverter(n, n.DatabaseName, n.MeasurementName)
	case common.ConvertPrometheus:
		dbConvert = prometheus.NewPrometheusConverter(n, n.DatabaseName, n.MeasurementName)
	}
	if dbConvert == nil {
		return nil, status.Error(codes.Unimplemented, "Not implement")
	}
	return dbConvert.GetWriteRequest()
}
