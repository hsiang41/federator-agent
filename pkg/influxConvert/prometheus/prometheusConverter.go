package prometheus

import (
	"time"
	"strings"

	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	v1alpha1common "github.com/containers-ai/api/common"
	"github.com/containers-ai/federatorai-agent/pkg"
	"github.com/golang/protobuf/ptypes"
)


type PrometheusMetrics []struct {
	Time   time.Time    `json:"Time"`
	Metric string       `json:"Metric"`
	Value  string       `json:"Value"`
}

type PMetrics struct {
	Label map[string] string
}

type prometheusConverter struct {
	TargetData      common.TargetData
	DatabaseName    string
	Measurement     string
}


func NewPrometheusConverter(targetdata common.TargetData, databaseName string, measurement string) *prometheusConverter {
	return &prometheusConverter{TargetData: targetdata, DatabaseName: databaseName, Measurement: measurement}
}

func addColumns(columns *[]string, columnName string) {
	for _, name := range *columns {
		if name == columnName {
			return
		}
	}
	*columns = append(*columns, columnName)
}

func (d *prometheusConverter) GetWriteRequest() (*v1alpha1.WriteRawdataRequest, error) {
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawData v1alpha1common.WriteRawdata
	var pDatas []*map[string]string
	var rawDatas []*v1alpha1common.WriteRawdata
	var metricColumnsTypes []v1alpha1common.ColumnType
	var metricDataTypes []v1alpha1common.DataType

	n := d.TargetData
	wrRawData.DatabaseType = v1alpha1common.DatabaseType_INFLUXDB

	souceData := n.GetSourceData().(*PrometheusMetrics)
	if len(*souceData) <= 0 {
		return nil, nil
	}

	for _, v := range *souceData {
		var fields map[string]string
		fields = make(map[string]string, 0)
		labels := v.Metric
		// metricsName := labels[:strings.Index(labels, "{")]
		labelEntry := strings.Split(labels[strings.Index(labels, "{")+1: len(labels) -1], ",")
		for _, v := range labelEntry {
			value := strings.Split(v, "=")
			fields[strings.TrimSpace(value[0])] = strings.Replace(value[1], "\"", "", -1)
			addColumns(&rawData.Columns, strings.TrimSpace(value[0]))
		}
		fields["time"] =  v.Time.Format(time.RFC3339Nano)
		// addColumns(&rawData.Columns, "time")
		fields["Value"] = v.Value
		addColumns(&rawData.Columns, "Value")
		pDatas = append(pDatas, &fields)
	}

	for _, c := range rawData.Columns {
		if c == "time" {
			continue
		}
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
	for _, g := range pDatas {
		var row v1alpha1common.Row
		tmValue, ok := (*g)["time"]
		if ok == false {
			continue
		}
		tm, _ := time.Parse(time.RFC3339Nano, tmValue)
		row.Time, _ = ptypes.TimestampProto(tm)
		for _, c := range rawData.Columns {
			value := ""
			v, ok := (*g)[c]
			if ok == true {
				value = v
			} else {
				value = ""
			}
			row.Values = append(row.Values, value)
		}
		if len(row.Values) > 0 {
			rawData.Rows = append(rawData.Rows, &row)
		}
	}
	rawData.Table = d.Measurement
	rawData.Database = d.DatabaseName
	rawData.ColumnTypes = metricColumnsTypes
	rawData.DataTypes = metricDataTypes
	rawDatas = append(rawDatas, &rawData)
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}
	return &wrRawData, nil
}
