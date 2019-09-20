package main

import (
	"io/ioutil"
	"net"
	"os"
	"errors"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	DataHubCommon "github.com/containers-ai/api/common"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/client/influx"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert"
	"github.com/containers-ai/federatorai-agent/pkg"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"github.com/containers-ai/federatorai-agent/pkg/client/prometheus"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert/prometheus"
	"github.com/containers-ai/federatorai-agent/pkg/influxConvert/influx"
	"github.com/containers-ai/federatorai-agent/pkg/datahub"
	"github.com/google/uuid"
)

var gCollector *collector

var FieldTypeMap = map[string] DataHubCommon.DataType {
	"float": DataHubCommon.DataType_DATATYPE_FLOAT64,
	"string": DataHubCommon.DataType_DATATYPE_STRING,
	"int": DataHubCommon.DataType_DATATYPE_FLOAT64,
}

type global struct {
	Interval    int     `mapstructure:"interval"`
}

type element struct {
	ElementType string  `mapstructure:"type"`
	ElementDefault string `mapstructure:"default"`
}

type Measurement struct {
	Name string `mapstructure:"name"`
	Expr string `mapstructure:"expr"`
	Tags []string `mapstructure:"tags"`
	LastSeconds string `mapstructure:"last_seconds"`
	Element map[string]element
}

type datasource struct {
	DataType    string `mapstructure:"datatype"`
	Address     string `mapstructure:"address"`
	Port        string `mapstructure:"port"`
	UserName    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	Measurements []*Measurement `mapstructure:"measurement"`
}

// Config defines configurations
type Config struct {
	Global      *global                 `mapstructure:"global"`
	Target      *datasource             `mapstructure:"target"`
	DataSource  map[string]datasource   `mapstructure:"datasource"`
}

type inputLib struct {
}

type collector struct {
	Config  *Config
	Logger  *logUtil.Scope
	Queue   *queue.Queue
}

func NewCollector(config *string, scope *logUtil.Scope) *collector {
	agentConf := &Config{}

	viper.SetEnvPrefix("InputLibPrometheusGPU")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	scope.Debugf("Load config path:(%p) %s", config, *config)

	viper.SetConfigFile(*config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library Prometheus GPU configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&agentConf)
	if err != nil {
		panic(errors.New("Unmarshal input library Prometheus GPU configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(agentConf, "", "  "); err == nil {
			scope.Debugf(fmt.Sprintf("Input library Prometheus GPU configuration: %s", string(transmitterConfBin)))
		} else {
			fmt.Printf("failed to display Prometheus GPU configuration")
		}
		return &collector{Config: agentConf, Logger: scope}
	}
	return nil
}

func (c *collector) HealthCheck() {
	var port string
	var address string
	if len(c.Config.DataSource) <= 0 {
		ioutil.WriteFile("status", []byte("0"), 0777)
		return
	}
	for _, v := range c.Config.DataSource {
		if len(v.Port) > 0 {
			port = v.Port
			address = v.Address
		} else {
			port = strings.Split(v.Address, ":")[0]
			address = strings.Split(v.Address, ":")[1]
		}
		if len(address) >= 7 && address[0:7] == "http://" {
			address = address[7:]
		}
		if len(address) >= 8 && address[0:8] == "https://" {
			address = address[8:]
		}
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", address, port))
		if err != nil {
			c.Logger.Errorf("Failed to connect datasource server: %s with %v", fmt.Sprintf("%s:%s", address, port), err)
			ioutil.WriteFile("status", []byte("0"), 0777)
			return
		}
		defer conn.Close()
	}
	err := ioutil.WriteFile("status", []byte("1"), 0777)
	if err != nil {
		c.Logger.Errorf("Failed to write agent running status with %v", err)
	}
}

func (c *collector) writeRawData (measurementName string, tags []string, fields *map[string]element, sourceData interface{}, convertType common.ConvertInt) error {
	var iFields [] *common.InfluxField
	for k, e := range *fields {
		iField := common.InfluxField{Name: k, Type: FieldTypeMap[strings.ToLower(e.ElementType)]}
		iFields = append(iFields, &iField)
	}
	iConvert := InfluxConvert.NewInflux(c.Config.Target.Database, measurementName, tags, iFields, sourceData, convertType)
	rawDatas, err := iConvert.GetWriteRequest()
	if err != nil {
		return err
	}

	dp := datahub.NewDataHubClient()
	dp.SetDataPipeAddress(c.Config.Target.Address)
	err = dp.WriteRawData(rawDatas)

	c.Logger.Debugf(utils.InterfaceToString(rawDatas))
	rawDatas = &v1alpha1.WriteRawdataRequest{}
	return err
}

func (c *collector) notifyEvent(eventLevel v1alpha1.EventLevel, message string, data string) error {
	hostName, _ := os.Hostname()
	dp := datahub.NewDataHubClient()
	dp.SetDataPipeAddress(c.Config.Target.Address)
	err := dp.SendNotifyEvent(uuid.New().String(), "",
		&v1alpha1.EventSource{Host: hostName, Component: "federatorai-agent"},
		v1alpha1.EventType_EVENT_TYPE_EMAIL_NOTIFICATION, eventLevel, nil, message, data)
	if err != nil {
		c.Logger.Errorf("Failed to send notification: %s %s", message, data)
	}
	return err
}

func (c *collector) Gather() error {
	for _, dv := range c.Config.DataSource {
		switch strings.ToLower(dv.DataType) {
		case "influx":
			for _, m := range dv.Measurements {
				var rawResponse influx.InfluxResp
				expr := fmt.Sprintf("%s where time > now() - %ss order by time asc", m.Expr, m.LastSeconds)
				// expr := fmt.Sprintf("%s order by time desc limit 10", m.Expr)
				c.Logger.Infof("expr: %s", expr)
				dbClient := ClientInflux.NewClientInflux(dv.Address, dv.Database, ClientInflux.MethodQuery, expr)
				result, err := dbClient.Execute()
				if err != nil {
					message := fmt.Sprintf("Failed to query %s with %v", m.Name, err)
					c.Logger.Errorf(message)
					c.notifyEvent(v1alpha1.EventLevel_EVENT_LEVEL_ERROR, message, "")
					continue
				}
				c.Logger.Debugf("result: %s", utils.InterfaceToString(result))
				err = json.Unmarshal([]byte(result), &rawResponse)
				if err != nil {
					c.Logger.Errorf("Unable to parse influx measurement %s, %v", m.Name, err)
					c.Logger.Errorf("result: %s", result)
					continue
				}
				c.Logger.Debugf("Start to write raw data %s, %s", dv.Database, utils.InterfaceToString(rawResponse))
				err = c.writeRawData(m.Name, m.Tags, &m.Element, &rawResponse, common.ConvertInflux)
				if err != nil {
					c.Logger.Errorf("Failed tp write %s result to raw data with %v", m.Name, err)
					continue
				}
				c.Logger.Infof("Succeed to write raw data %s", m.Name)
			}
		case "prometheus":
			for _, m := range dv.Measurements {
				var rawResponse prometheus.PrometheusMetrics
				var step string
				tm := utils.TimeRange{}
				tm.EndTime = time.Now()
				if len(m.LastSeconds) > 0 {
					step = fmt.Sprintf("%ss", m.LastSeconds)
				} else {
					step = fmt.Sprintf("%ds", c.Config.Global.Interval)
				}
				if step == "1s" {
					tm.StartTime = tm.EndTime
					tm.Step, _ = time.ParseDuration(fmt.Sprintf("%s", step))
				} else {
					td, _ := time.ParseDuration(fmt.Sprintf("-%s", step))
					tm.StartTime = tm.EndTime.Add(td)
					tm.Step, _ = time.ParseDuration(fmt.Sprintf("%s", step))
				}
				expr := m.Expr
				dbClient := ClientPrometheus.NewClientPrometheus(dv.Address, ClientPrometheus.MethodQueryRange, expr, &tm)
				result, err := dbClient.Execute()
				if err != nil {
					message := fmt.Sprintf("Failed to query %s with %v", m.Name, err)
					c.Logger.Errorf(message)
					c.notifyEvent(v1alpha1.EventLevel_EVENT_LEVEL_ERROR, message, "")
					continue
				}
				if len(result) == 0 || strings.ToLower(result) == "none" {
					continue
				}
				err = json.Unmarshal([]byte(result), &rawResponse)
				if err != nil {
					c.Logger.Errorf("Unable to parse prometheus %s, %v", m.Name, err)
					c.Logger.Errorf("result: %s", result)
					continue
				}
				c.Logger.Debugf("Start to write raw data %s, %s", dv.Database, utils.InterfaceToString(rawResponse))
				err = c.writeRawData(m.Name, m.Tags, &m.Element, &rawResponse, common.ConvertPrometheus)
				if err != nil {
					// c.Logger.Errorf("Failed to write %s result to raw data with %v", m.Name, err)
					continue
				}
				c.Logger.Infof("Succeed to write raw data %s", m.Name)
			}
		}
	}
	return nil
}

func (i inputLib) Gather() error {
	return gCollector.Gather()
}

func (i inputLib) LoadConfig(config *string, scope *logUtil.Scope) error {
	gCollector = NewCollector(config, scope)
	return nil
}

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gCollector.Queue = agentQueue
}


func (i inputLib) Close() {
	gCollector = &collector{}
}

var InputLib inputLib
