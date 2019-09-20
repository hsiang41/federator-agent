package influx

import (
	"fmt"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	v1alpha1common "github.com/containers-ai/api/common"
	"github.com/containers-ai/federatorai-agent/pkg")

type influxConverter struct {
	TargetData      common.TargetData
	DatabaseName    string
	Measurement     string
}

/*
type InfluxResp struct {
	Results []struct {
		Series []struct {
			Name    string          `json:"name"`
			Columns []string        `json:"columns"`
			Values  [][]interface{} `json:"values"`
		} `json:"Series"`
		Messages interface{} `json:"Messages"`
	} `json:"Results"`
}
*/

type InfluxResp []struct {
	Series []struct {
		Name    string          `json:"name"`
		Columns []string        `json:"columns"`
		Values  [][]interface{} `json:"values"`
	} `json:"Series"`
	Messages interface{} `json:"Messages"`
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

func NewInfluxConverter(targetdata common.TargetData, databaseName string, measurement string) *influxConverter {
	return &influxConverter{TargetData: targetdata, DatabaseName: databaseName, Measurement: measurement}
}

func (d *influxConverter) GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error) {
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*v1alpha1common.WriteRawdata

	n := d.TargetData
	wrRawData.DatabaseType = v1alpha1common.DatabaseType_INFLUXDB
	if len(*n.GetSourceData().(*InfluxResp)) <= 0 {
		return nil, status.Error(codes.NotFound, "Not data parse")
	}
	souceData := n.GetSourceData().(*InfluxResp)
	for _, v := range (*souceData)[0].Series {
		var rawData v1alpha1common.WriteRawdata
		var metricColumnsTypes []v1alpha1common.ColumnType
		var metricDataTypes []v1alpha1common.DataType
		if len(v.Columns) <= 1 {
			continue
		}
		rawData.Columns = v.Columns[1:]
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
		for _, g := range v.Values {
			var row v1alpha1common.Row
			for i, gv := range g {
				var value string
				if i == 0 {
					tm, _ := time.Parse(time.RFC3339Nano, gv.(string))
					row.Time, _ = ptypes.TimestampProto(tm)
					continue
				}
				rv := reflect.ValueOf(gv)
				switch rv.Kind() {
				case reflect.String:
					value = gv.(string)
				case reflect.Int, reflect.Int64:
					value = fmt.Sprintf("%d", gv.(int64))
				case reflect.Float32, reflect.Float64:
					value = fmt.Sprintf("%f", gv.(float64))
				default:
					value = ""
				}
				row.Values = append(row.Values, value)
			}
			rawData.Rows = append(rawData.Rows, &row)
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
