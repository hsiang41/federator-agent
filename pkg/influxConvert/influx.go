package InfluxConvert

import (
	"fmt"
	"strings"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/common"
)

type InfluxField struct {
	Name string
	Type common.DataType
	Default interface{}
}

type Influx struct {
	DatabaseName    string
	MeasurementName string
	TagKeys         []string
	FieldKeys       []*InfluxField
	SourceData      interface{}
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

func NewInflux(databaseName string, measurementName string, tags []string, fields []*InfluxField, sourceData interface{}) *Influx {
	return &Influx{
		DatabaseName:    databaseName,
		MeasurementName: measurementName,
		TagKeys:         tags, FieldKeys: fields, SourceData: sourceData}
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


func (n *Influx)GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error) {
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	func () {
		defer func() {
			fmt.Println("defer here")
		}()
	}()

	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	souceData := n.SourceData.(*v1alpha1.ReadRawdataResponse).GetRawdata()

    for _, v := range souceData {
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