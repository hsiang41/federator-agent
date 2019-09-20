package influx_fedemeter

import (
	"encoding/json"
	"fmt"
	"strings"
	"reflect"
	"strconv"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/common"
	"github.com/containers-ai/federatorai-agent/pkg/fedemeter"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type EnumMeasurementID int32

var EnumMeasurementName = map[EnumMeasurementID]string {
	0: "calculate_instance",
	1: "calculate_storage",
	2: "recommendation_jeri",
	3: "prediction_cost_namespace",
	4: "prediction_cost_app",
	5: "history_cost_namespace",
	6: "history_cost_app"}

var MeasurementColumns map[EnumMeasurementID][]string

func init() {
	MeasurementColumns = make(map[EnumMeasurementID][]string, 0)
	MeasurementColumns[0] = []string{"starttime", "provider", "nodename", "unit", "instancetype", "cpu", "memory", "totalcost", "displayname", "instancenum", "region", "description", "cost", "granularity"}
	MeasurementColumns[1] = []string{"starttime", "nodename", "volumetype", "unit", "storagesize", "description", "cost", "displayname", "granularity"}
	MeasurementColumns[2] = []string{"starttime", "resourcename", "provider", "granularity", "region", "instancetype", "totalcost", "masternum", "workernum", "masterstoragesize", "workerstoragesize", "ondemandnum", "revservedinstances",
		"displayname", "master_ri_num", "worker_ri_num", "master_ondemand_num", "worker_ondemand_num"}
	MeasurementColumns[3] = []string{"clustername", "provider", "namespacesname", "workloadcost", "costpercentage", "granularity"}
	MeasurementColumns[4] = []string{"clustername", "provider", "namespacesname", "appname", "workloadcost", "costpercentage", "granularity"}
	MeasurementColumns[5] = []string{"clustername", "provider", "namespacesname", "workloadcost", "costpercentage", "granularity"}
	MeasurementColumns[6] = []string{"clustername", "provider", "namespacesname", "appname", "workloadcost", "costpercentage", "granularity"}
}

type InfluxField struct {
	Name string
	Type common.DataType
	Default interface{}
}

type InfluxMeasurement struct {
	DatabaseName      string
	MeasurementID     EnumMeasurementID
	TagKeys           []string
	FieldKeys         []*InfluxField
	SourceData        interface{}
	Granularity       int64
	ReservedInstances bool
}

type recommendatioinJeri struct {
	StartTime         int64
	ResourceName      string
	Provider          string
	Granularity       int64
	Region            string
	InstanceType      string
	TotalCost         float64
	MasterNum         int
	WorkerNum         int
	MasterStorageSize float64
	WorkerStorageSize float64
	OndemandNum       int
	ReservedInstances int
	MasterRiNum       int
	WorkerRiNum       int
	MasterOndemandNum int
	WorkerOndemandNum int
	DisplayName       string
}

type costNamespace struct {
	ClusterName     string
	Provider        string
	NamespaceName   string
	WorkloadCost    float64
	CostPercentage  float64
	Granularity     int64
	Time            int64
}

type costApp struct {
	ClusterName     string
	Provider        string
	NamespaceName   string
	AppName         string
	WorkloadCost    float64
	CostPercentage  float64
	Granularity     int64
	Time            int64
}


func NewInfluxMeasurement(databaseName string, measurementID EnumMeasurementID, tags []string, fields []*InfluxField, sourceData interface{}, granularity int64, revservedinstances bool) *InfluxMeasurement {
	return &InfluxMeasurement{
		DatabaseName:  databaseName,
		MeasurementID: measurementID,
		TagKeys:       tags, FieldKeys: fields, SourceData: sourceData, Granularity: granularity, ReservedInstances: revservedinstances}
}

func (n *InfluxMeasurement)isTagKey(keyName string) bool {
	if n.TagKeys == nil {
		return false
	}
	for _, n := range n.TagKeys {
		if n == strings.ToLower(keyName) {
			return true
		}
	}
	return false
}

func (n *InfluxMeasurement)getFieldKey(keyName string) *InfluxField {
	for _, k := range MeasurementColumns[n.MeasurementID] {
		if strings.ToLower(keyName) == k {
			for _, n := range n.FieldKeys {
				if n.Name == strings.ToLower(keyName) {
					return n
				}
			}
			return &InfluxField{strings.ToLower(keyName), common.DataType_DATATYPE_STRING, nil}
		}
	}
	return nil
}

func (n *InfluxMeasurement) generateCalculateInstance(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	n.TagKeys = []string {"provide", "nodename", "unit", "ganularity"}
	fieldTotalCost := &InfluxField{"totalcost", common.DataType_DATATYPE_FLOAT64, nil}
	fieldCost := &InfluxField{"cost", common.DataType_DATATYPE_FLOAT64, nil}
	if n.FieldKeys == nil {
		n.FieldKeys = make([]*InfluxField, 0)
	}
	n.FieldKeys = append(n.FieldKeys, fieldTotalCost)
	n.FieldKeys = append(n.FieldKeys, fieldCost)

	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	var rawData common.WriteRawdata

	startTime := starttime
	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	souceData := n.SourceData.(*fedemeter.FedCalculatorResp)

	// Set DatabaseName and MeasurementName
	rawData.Table = EnumMeasurementName[n.MeasurementID]
	rawData.Database = n.DatabaseName
	initialMeasurementField := 0
	for _, v := range souceData.Calculator {
		for k, instances := range v {
			provider := k
			for _, iv := range instances {
				totalCoat := iv.Instances.Cost
				for _, s := range iv.Storage {
					totalCoat += s.Cost
				}
				totalCoat += iv.Gpu.Cost

				rk := reflect.TypeOf(&iv.Instances).Elem()
				rv := reflect.ValueOf(&iv.Instances).Elem()

				// Config tags, fields
				if initialMeasurementField == 0 {
					var metricColumnsTypes []common.ColumnType
					var metricDataTypes []common.DataType
					for i := 0; i < rk.NumField(); i++ {
						// Set Influx tags
						if n.isTagKey(rk.Field(i).Name) == true {
							metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
							metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
							rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
							continue
						}
						// Set Influx fields
						metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
						iField := n.getFieldKey(rk.Field(i).Name)
						if iField == nil {
							metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
						} else {
							metricDataTypes = append(metricDataTypes, iField.Type)
						}
						rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
					}
					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
					rawData.Columns = append(rawData.Columns, "provider")

					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
					fieldTotalCost := InfluxField{"totalcost", common.DataType_DATATYPE_FLOAT64, nil}
					metricDataTypes = append(metricDataTypes, fieldTotalCost.Type)
					rawData.Columns = append(rawData.Columns, strings.ToLower(fieldTotalCost.Name))

					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
					fieldStartTime := InfluxField{"starttime", common.DataType_DATATYPE_INT64, nil}
					metricDataTypes = append(metricDataTypes, fieldStartTime.Type)
					rawData.Columns = append(rawData.Columns, strings.ToLower(fieldStartTime.Name))

					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
					fieldGranularity := InfluxField{"granularity", common.DataType_DATATYPE_INT64, nil}
					metricDataTypes = append(metricDataTypes, fieldStartTime.Type)
					rawData.Columns = append(rawData.Columns, strings.ToLower(fieldGranularity.Name))

					rawData.ColumnTypes = metricColumnsTypes
					rawData.DataTypes = metricDataTypes
					initialMeasurementField = 1
				}

				// Set raw data
				var rawValus common.Row
				rawValus.Time = &timestamp.Timestamp{Seconds: startTime.Seconds, Nanos: startTime.Nanos}
				for i := 0; i< rk.NumField(); i++ {
					var value string
					switch rv.FieldByName(rk.Field(i).Name).Kind() {
					case reflect.Int, reflect.Int64:
						value = fmt.Sprintf("%d", rv.FieldByName(rk.Field(i).Name).Int())
					case reflect.Float64, reflect.Float32:
						value = fmt.Sprintf("%f", rv.FieldByName(rk.Field(i).Name).Float())
					default:
						value = rv.FieldByName(rk.Field(i).Name).String()
					}
					rawValus.Values = append(rawValus.Values, value)
				}
				// append total cost
				rawValus.Values = append(rawValus.Values, provider)
				rawValus.Values = append(rawValus.Values, fmt.Sprintf("%f", totalCoat))
				rawValus.Values = append(rawValus.Values, fmt.Sprintf("%d", startTime.Seconds))
				rawValus.Values = append(rawValus.Values, fmt.Sprintf("%d", n.Granularity))
				rawData.Rows = append(rawData.Rows, &rawValus)
			}
		}
	}
	if len(rawData.Rows) > 0 {
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}
	return &wrRawData, nil
}

func (n *InfluxMeasurement) generateCalculateStorage(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	n.TagKeys = []string {"nodename", "provider", "unit", "granularity"}
	fieldCost := &InfluxField{"cost", common.DataType_DATATYPE_FLOAT64, nil}
	if n.FieldKeys == nil {
		n.FieldKeys = make([]*InfluxField, 0)
	}
	n.FieldKeys = append(n.FieldKeys, fieldCost)

	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	var rawData common.WriteRawdata

	startTime := starttime
	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	souceData := n.SourceData.(*fedemeter.FedCalculatorResp)

	// Set DatabaseName and MeasurementName
	rawData.Table = EnumMeasurementName[n.MeasurementID]
	rawData.Database = n.DatabaseName
	initialMeasurementField := 0
	for _, v := range souceData.Calculator {
		for k, instances := range v {
			provider := k
			for _, iv := range instances {
				nodeName := iv.Instances.Nodename
				for _, s := range iv.Storage {
					rk := reflect.TypeOf(&s).Elem()
					rv := reflect.ValueOf(&s).Elem()
					if initialMeasurementField == 0 {
						var metricColumnsTypes []common.ColumnType
						var metricDataTypes []common.DataType
						for i := 0; i < rk.NumField(); i++ {
							// Set Influx tags
							if n.isTagKey(rk.Field(i).Name) == true {
								metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
								metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
								rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
								continue
							}
							// Set Influx fields
							metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
							iField := n.getFieldKey(rk.Field(i).Name)
							if iField == nil {
								metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
							} else {
								metricDataTypes = append(metricDataTypes, iField.Type)
							}
							rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
						}
						metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
						metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
						rawData.Columns = append(rawData.Columns, "nodename")

						metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
						metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
						rawData.Columns = append(rawData.Columns, "provider")

						metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
						fieldStartTime := InfluxField{"starttime", common.DataType_DATATYPE_INT64, nil}
						metricDataTypes = append(metricDataTypes, fieldStartTime.Type)
						rawData.Columns = append(rawData.Columns, strings.ToLower(fieldStartTime.Name))

						metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
						fieldGranularity := InfluxField{"granularity", common.DataType_DATATYPE_INT64, nil}
						metricDataTypes = append(metricDataTypes, fieldStartTime.Type)
						rawData.Columns = append(rawData.Columns, strings.ToLower(fieldGranularity.Name))

						rawData.ColumnTypes = metricColumnsTypes
						rawData.DataTypes = metricDataTypes
						initialMeasurementField = 1
					}
					// Set raw data
					var rawValus common.Row
					rawValus.Time = &timestamp.Timestamp{Seconds: startTime.Seconds, Nanos: startTime.Nanos}
					for i := 0; i< rk.NumField(); i++ {
						var value string
						switch rv.FieldByName(rk.Field(i).Name).Kind() {
						case reflect.Int, reflect.Int64:
							value = fmt.Sprintf("%d", rv.FieldByName(rk.Field(i).Name).Int())
						case reflect.Float64, reflect.Float32:
							value = fmt.Sprintf("%f", rv.FieldByName(rk.Field(i).Name).Float())
						default:
							value = rv.FieldByName(rk.Field(i).Name).String()
						}
						rawValus.Values = append(rawValus.Values, value)
					}
					// append total cost
					rawValus.Values = append(rawValus.Values, nodeName)
					rawValus.Values = append(rawValus.Values, provider)
					rawValus.Values = append(rawValus.Values, fmt.Sprintf("%d", startTime.Seconds))
					rawValus.Values = append(rawValus.Values, fmt.Sprintf("%d", n.Granularity))
					rawData.Rows = append(rawData.Rows, &rawValus)
				}
			}
		}
	}
	if len(rawData.Rows) > 0 {
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}
	return &wrRawData, nil
}

func (n *InfluxMeasurement)generateCalculateRecommendation(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	var rcDatas []*recommendatioinJeri
	n.TagKeys = []string {"provider", "instancetype", "granularity", "reservedinstances"}
	fieldCost := &InfluxField{"cost", common.DataType_DATATYPE_FLOAT64, nil}
	if n.FieldKeys == nil {
		n.FieldKeys = make([]*InfluxField, 0)
	}
	n.FieldKeys = append(n.FieldKeys, fieldCost)
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"totalcost", common.DataType_DATATYPE_FLOAT64, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"masterstoragesize", common.DataType_DATATYPE_FLOAT64, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"workerstoragesize", common.DataType_DATATYPE_FLOAT64, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"workernum", common.DataType_DATATYPE_INT, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"masternum", common.DataType_DATATYPE_INT, nil})
	souceData := n.SourceData.(*fedemeter.FedRecommendationJeriResp)

	for sk, sv := range souceData.Resource {
		resourceName := sk
		for _, pK := range sv {
			for _, pKI := range pK {
				for provider, inst := range pKI {
					var region string
					for k, v := range inst {
						var instanceType string
						var riInstance []fedemeter.FedJeriInstance
						switch k {
						case "region":
							err := json.Unmarshal(v, &region)
							if err != nil {
								region = ""
							}
						default:
							instanceType = k
							riInstance = make([]fedemeter.FedJeriInstance, 0)
							err := json.Unmarshal(v, &riInstance)
							if err != nil {
								fmt.Println(err)
								riInstance = nil
							}
							if riInstance != nil && len(riInstance) > 0 {
								var rcData recommendatioinJeri
								rcData.StartTime = starttime.Seconds
								lstData := riInstance[len(riInstance) - 1]
								totalcost := lstData.AccCost
								rcData.TotalCost = totalcost
								rcData.MasterNum = lstData.MasterNum
								if n.ReservedInstances == true {
									rcData.ReservedInstances = 1
								} else {
									rcData.ReservedInstances = 0
								}
								rcData.Granularity = n.Granularity
								rcData.InstanceType = instanceType
								rcData.OndemandNum = lstData.OndemandNum
								rcData.MasterStorageSize = lstData.MasterStorageSize
								rcData.WorkerStorageSize = lstData.WorkerStorageSize
								rcData.DisplayName = lstData.Displayname
								rcData.MasterRiNum = lstData.MasterRiNum
								rcData.WorkerRiNum = lstData.WorkerRiNum
								rcData.MasterOndemandNum = lstData.MasterOndemandNum
								rcData.WorkerOndemandNum = lstData.WorkerOndemandNum
								rcData.Provider = provider
								rcData.WorkerNum = lstData.WorkerNum
								rcData.ResourceName = resourceName
								rcData.Region = region
								rcDatas = append(rcDatas, &rcData)
							}
						}
					}
				}
			}
		}
	}

	// Reformat to influx data fields
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	var rawData common.WriteRawdata

	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	// Set DatabaseName and MeasurementName
	rawData.Table = EnumMeasurementName[n.MeasurementID]
	rawData.Database = n.DatabaseName
	initialMeasurementField := 0
	for _, d := range rcDatas {
		rk := reflect.TypeOf(d).Elem()
		rv := reflect.ValueOf(d).Elem()
		if initialMeasurementField == 0 {
			var metricColumnsTypes []common.ColumnType
			var metricDataTypes []common.DataType
			for i := 0; i < rk.NumField(); i++ {
				// Set Influx tags
				if n.isTagKey(rk.Field(i).Name) == true {
					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
					rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
					continue
				}
				// Set Influx fields
				metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
				iField := n.getFieldKey(rk.Field(i).Name)
				if iField == nil {
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
				} else {
					metricDataTypes = append(metricDataTypes, iField.Type)
				}
				rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
			}

			rawData.ColumnTypes = metricColumnsTypes
			rawData.DataTypes = metricDataTypes
			initialMeasurementField = 1
		}
		// Set raw data
		var rawValus common.Row
		rawValus.Time = &timestamp.Timestamp{Seconds: starttime.Seconds, Nanos: starttime.Nanos}
		for i := 0; i< rk.NumField(); i++ {
			var value string
			switch rv.FieldByName(rk.Field(i).Name).Kind() {
			case reflect.Int, reflect.Int64:
				value = fmt.Sprintf("%d", rv.FieldByName(rk.Field(i).Name).Int())
			case reflect.Float64, reflect.Float32:
				value = fmt.Sprintf("%f", rv.FieldByName(rk.Field(i).Name).Float())
			default:
				value = rv.FieldByName(rk.Field(i).Name).String()
			}
			rawValus.Values = append(rawValus.Values, value)
		}
		// append total cost
		rawData.Rows = append(rawData.Rows, &rawValus)
	}
	if len(rawData.Rows) > 0 {
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}
	return &wrRawData, nil
}

func (n *InfluxMeasurement) generateCostNamespace(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	var costNamespaces []*costNamespace
	n.TagKeys = []string {"clustername", "provider", "namespacename", "granularity"}
	if n.FieldKeys == nil {
		n.FieldKeys = make([]*InfluxField, 0)
	}
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"workloadcost", common.DataType_DATATYPE_FLOAT64, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"costpercentage", common.DataType_DATATYPE_FLOAT64, nil})
	if n.SourceData == nil {
		return nil, status.Error(codes.Unavailable, "Unable to parse data source")
	}
	souceData := n.SourceData.(*fedemeter.FedCostMetricResp)

	clusterName := souceData.Cluster.Clustername
	for _, p := range souceData.Cluster.Providers {
		provider := p.Providername
		for _, namespace := range p.Namespace {
			for _, cost := range namespace.Costs {
				namespacesName := namespace.Namespacename
				workloadCost, _ := strconv.ParseFloat(cost.Workloadcost, 64)
				costPercentage, _ := strconv.ParseFloat(cost.Costpercentage, 64)
				costNamespace := &costNamespace{ClusterName: clusterName, Provider: provider, NamespaceName: namespacesName, WorkloadCost: workloadCost, CostPercentage: costPercentage, Time: cost.Timestampe * n.Granularity, Granularity: n.Granularity}
				costNamespaces = append(costNamespaces, costNamespace)
			}
		}
	}

	if costNamespaces == nil || len(costNamespaces) == 0 {
		return nil, nil
	}


	// Reformat to influx data fields
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	var rawData common.WriteRawdata

	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	// Set DatabaseName and MeasurementName
	rawData.Table = EnumMeasurementName[n.MeasurementID]
	rawData.Database = n.DatabaseName
	initialMeasurementField := 0
	for _, d := range costNamespaces {
		rk := reflect.TypeOf(d).Elem()
		rv := reflect.ValueOf(d).Elem()
		if initialMeasurementField == 0 {
			var metricColumnsTypes []common.ColumnType
			var metricDataTypes []common.DataType
			for i := 0; i < rk.NumField(); i++ {
				// Set Influx tags
				if rk.Field(i).Name == "Time" {
					continue
				}
				if n.isTagKey(rk.Field(i).Name) == true {
					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
					rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
					continue
				}
				// Set Influx fields
				metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
				iField := n.getFieldKey(rk.Field(i).Name)
				if iField == nil {
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
				} else {
					metricDataTypes = append(metricDataTypes, iField.Type)
				}
				rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
			}

			rawData.ColumnTypes = metricColumnsTypes
			rawData.DataTypes = metricDataTypes
			initialMeasurementField = 1
		}
		// Set raw data
		var rawValus common.Row
		for i := 0; i< rk.NumField(); i++ {
			var value string
			if rk.Field(i).Name == "Time" {
				rawValus.Time = &timestamp.Timestamp{Seconds: rv.FieldByName(rk.Field(i).Name).Int()}
				continue
			}
			switch rv.FieldByName(rk.Field(i).Name).Kind() {
			case reflect.Int, reflect.Int64:
				value = fmt.Sprintf("%d", rv.FieldByName(rk.Field(i).Name).Int())
			case reflect.Float64, reflect.Float32:
				value = fmt.Sprintf("%f", rv.FieldByName(rk.Field(i).Name).Float())
			default:
				value = rv.FieldByName(rk.Field(i).Name).String()
			}
			rawValus.Values = append(rawValus.Values, value)
		}
		// append total cost
		rawData.Rows = append(rawData.Rows, &rawValus)
	}
	if len(rawData.Rows) > 0 {
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}

	return &wrRawData, nil
}

func (n *InfluxMeasurement) generateCostApp(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	var costApps []*costApp
	n.TagKeys = []string {"clustername", "provider", "namespacename", "granularity"}
	if n.FieldKeys == nil {
		n.FieldKeys = make([]*InfluxField, 0)
	}
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"workloadcost", common.DataType_DATATYPE_FLOAT64, nil})
	n.FieldKeys = append(n.FieldKeys, &InfluxField{"costpercentage", common.DataType_DATATYPE_FLOAT64, nil})
	if n.SourceData == nil {
		return nil, status.Error(codes.Unavailable, "Unable to parse data source")
	}
	souceData := n.SourceData.(*fedemeter.FedCostMetricResp)

	clusterName := souceData.Cluster.Clustername
	for _, p := range souceData.Cluster.Providers {
		provider := p.Providername
		for _, namespace := range p.Namespace {
			for _, app := range namespace.Apps {
				appName := app.Appname
				for _, cost := range app.Costs {
					namespacesName := namespace.Namespacename
					workloadCost, _ := strconv.ParseFloat(cost.Workloadcost, 64)
					costPercentage, _ := strconv.ParseFloat(cost.Costpercentage, 64)
					costApp := &costApp{ClusterName: clusterName, Provider: provider, NamespaceName: namespacesName, WorkloadCost: workloadCost, CostPercentage: costPercentage, Time: cost.Timestampe * n.Granularity, AppName: appName, Granularity: n.Granularity}
					costApps = append(costApps, costApp)
				}
			}
		}
	}

	if costApps == nil || len(costApps) == 0 {
		return nil, nil
	}

	// Reformat to influx data fields
	var wrRawData v1alpha1.WriteRawdataRequest
	var rawDatas []*common.WriteRawdata
	var rawData common.WriteRawdata

	wrRawData.DatabaseType = common.DatabaseType_INFLUXDB
	// Set DatabaseName and MeasurementName
	rawData.Table = EnumMeasurementName[n.MeasurementID]
	rawData.Database = n.DatabaseName
	initialMeasurementField := 0
	for _, d := range costApps {
		rk := reflect.TypeOf(d).Elem()
		rv := reflect.ValueOf(d).Elem()
		if initialMeasurementField == 0 {
			var metricColumnsTypes []common.ColumnType
			var metricDataTypes []common.DataType
			for i := 0; i < rk.NumField(); i++ {
				// Set Influx tags
				if rk.Field(i).Name == "Time" {
					continue
				}
				if n.isTagKey(rk.Field(i).Name) == true {
					metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_TAG)
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
					rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
					continue
				}
				// Set Influx fields
				metricColumnsTypes = append(metricColumnsTypes, common.ColumnType_COLUMNTYPE_FIELD)
				iField := n.getFieldKey(rk.Field(i).Name)
				if iField == nil {
					metricDataTypes = append(metricDataTypes, common.DataType_DATATYPE_STRING)
				} else {
					metricDataTypes = append(metricDataTypes, iField.Type)
				}
				rawData.Columns = append(rawData.Columns, strings.ToLower(rk.Field(i).Name))
			}

			rawData.ColumnTypes = metricColumnsTypes
			rawData.DataTypes = metricDataTypes
			initialMeasurementField = 1
		}
		// Set raw data
		var rawValus common.Row
		for i := 0; i< rk.NumField(); i++ {
			var value string
			if rk.Field(i).Name == "Time" {
				rawValus.Time = &timestamp.Timestamp{Seconds: rv.FieldByName(rk.Field(i).Name).Int()}
				continue
			}
			switch rv.FieldByName(rk.Field(i).Name).Kind() {
			case reflect.Int, reflect.Int64:
				value = fmt.Sprintf("%d", rv.FieldByName(rk.Field(i).Name).Int())
			case reflect.Float64, reflect.Float32:
				value = fmt.Sprintf("%f", rv.FieldByName(rk.Field(i).Name).Float())
			default:
				value = rv.FieldByName(rk.Field(i).Name).String()
			}
			rawValus.Values = append(rawValus.Values, value)
		}
		// append total cost
		rawData.Rows = append(rawData.Rows, &rawValus)
	}
	if len(rawData.Rows) > 0 {
		rawDatas = append(rawDatas, &rawData)
	}
	if len(rawDatas) > 0 {
		wrRawData.Rawdata = rawDatas
	}

	return &wrRawData, nil
}

func (n *InfluxMeasurement)GetWriteRequest(starttime *timestamp.Timestamp) (*v1alpha1.WriteRawdataRequest, error) {
	switch n.MeasurementID {
	case 0:
		return n.generateCalculateInstance(starttime)
	case 1:
		return n.generateCalculateStorage(starttime)
	case 2:
		return n.generateCalculateRecommendation(starttime)
	case 3, 5:
		return n.generateCostNamespace(starttime)
	case 4, 6:
		return n.generateCostApp(starttime)
	}
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Unimplment measurement id %d", n.MeasurementID))
}
