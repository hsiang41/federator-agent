package datahub

import (
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	v1alpha1common "github.com/containers-ai/api/common"
	"github.com/containers-ai/federatorai-agent/pkg"
)

type datahubConverter struct {
	TargetData      common.TargetData
	DatabaseName    string
	Measurement     string
}

/*
func StringToDataType(strDataType string) v1alpha1common.DataType {
	strDataType = strings.ToLower(strDataType)
	switch strDataType {
	case "int":
		return v1alpha1common.DataType_DATATYPE_INT
	case "float":
		return v1alpha1common.DataType_DATATYPE_FLOAT64
	default:
		return v1alpha1common.DataType_DATATYPE_STRING
	}
}
*/

func NewDatahubConverter(targetdata common.TargetData, databaseName string, measurement string) *datahubConverter {
	return &datahubConverter{TargetData: targetdata, DatabaseName: databaseName, Measurement: measurement}
}

func (d *datahubConverter) GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error) {
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*v1alpha1common.WriteRawdata

	n := d.TargetData
	wrRawData.DatabaseType = v1alpha1common.DatabaseType_INFLUXDB
	souceData := n.GetSourceData().(*v1alpha1.ReadRawdataResponse).GetRawdata()

	for _, v := range souceData {
		var rawData v1alpha1common.WriteRawdata
		var metricColumnsTypes []v1alpha1common.ColumnType
		var metricDataTypes []v1alpha1common.DataType
		rawData.Columns = v.Columns
		for _, c := range rawData.Columns {
			if n.IsTagKey(c) == true {
				metricColumnsTypes = append(metricColumnsTypes, v1alpha1common.ColumnType_COLUMNTYPE_TAG)
				metricDataTypes = append(metricDataTypes, v1alpha1common.DataType_DATATYPE_STRING)
			} else {
				metricColumnsTypes = append(metricColumnsTypes, v1alpha1common.ColumnType_COLUMNTYPE_FIELD)
				iField := n.GetFieldKey(c)
				if iField == nil {
					metricDataTypes = append(metricDataTypes, v1alpha1common.DataType_DATATYPE_STRING)
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
		rawData.Table = d.Measurement
		rawData.Database = d.DatabaseName
		rawData.ColumnTypes = metricColumnsTypes
		rawData.DataTypes = metricDataTypes
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}
	return &wrRawData, nil
}