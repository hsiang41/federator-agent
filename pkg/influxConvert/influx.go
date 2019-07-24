package InfluxConvert

import (
	"github.com/containers-ai/api/datapipe/rawdata"
	"github.com/containers-ai/api/common"
	"strings"
)

type InfluxField struct {
	Name string
	Type common.DataType
	Default interface{}
}

type Influx struct {
	DatabaseName string
	MeasurementName string
	TagKeys []string
	FieldKeys []*InfluxField
	sourceData *rawdata.ReadRawdataResponse
}

func StringToDataType(strDataType string) common.DataType {
	strDataType = strings.ToLower(strDataType)
	switch strDataType {
	case "int":
		return common.DataType_DATATYPE_INT
	case "float":
		return common.DataType_DATATYPE_FLOAT64
	default:
		return common.DataType_DATATYPE_STRING
	}
}

func NewInflux(databaseName string, measurementName string, tags []string, fields []*InfluxField, sourceData *rawdata.ReadRawdataResponse) *Influx {
	return &Influx{
		DatabaseName: databaseName,
		MeasurementName: measurementName,
		TagKeys: tags, FieldKeys: fields, sourceData: sourceData}
}

func (n *Influx)isTagKey(keyName string) bool {
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

func (n *Influx)getFieldKey(keyName string) *InfluxField {
	for _, n := range n.FieldKeys {
		if n.Name == keyName {
			return n
		}
	}
	return nil
}

func (n *Influx)GetWriteRequest() (*rawdata.WriteRawdataRequest, error) {
	var wrRawData rawdata.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
    for _, v := range n.sourceData.GetRawdata() {
		var rawData common.WriteRawdata
	    var metricColumnsTypes []common.ColumnType
	    var metricDataTypes []common.DataType
		rawData.Columns = v.Columns
		for _, c := range rawData.Columns {
			if n.isTagKey(c) == true {
				metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
				metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
			} else {
				metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
				iField := n.getFieldKey(c)
				if iField == nil {
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
				} else {
					metricDataTypes = append(metricDataTypes, iField.Type)
				}
			}
		}
		for _, g := range v.GetGroups() {
			if len(rawData.Rows) == 0 {
				rawData.Rows = g.Rows
			} else {
				rawData.Rows = append(rawData.Rows, g.Rows...)
			}
		}
		rawData.Table = n.MeasurementName
		rawData.Database = n.DatabaseName
		rawData.ColumnTypes = metricColumnsTypes
		rawData.DataTypes = metricDataTypes
		rawDatas = append(rawDatas, &rawData)
    }
    if len(rawDatas) > 0 {
    	wrRawData.Rawdata = rawDatas
    }
	return &wrRawData, nil
}