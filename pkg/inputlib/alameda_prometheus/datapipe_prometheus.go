package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"strings"
	"time"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"github.com/containers-ai/federatorai-agent/pkg/prometheus"
	"github.com/golang/protobuf/ptypes"
	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"

	"github.com/containers-ai/federatorai-agent/pkg/influxConvert"
)

type global struct {
	PrometheusScrapeStepSeconds int `mapstructure:"prometheus_scrape_step_seconds"`
	PrometheusScrapeFrequencySeconds int `mapstructure:"prometheus_scrape_frequency_seconds"`
	TargetDatabase string `mapstructure:"target_database"`
	TargetAddress string `mapstructure:"target_address"`
	TargetUser string `mapstructure:"target_user"`
	TargetPassword string `mapstructure:"target_password"`
}

type element struct {
	ElementType string `mapstructure:"type"`
	ElementDefault string `mapstructure:"default"`
}

type measurement struct {
	Name string `mapstructure:"name"`
	Expr string `mapstructure:"expr"`
	Tags []string `mapstructure:"tags"`
	Element map[string]element
}

// Config defines configurations
type Config struct {
	Global  *global `mapstructure:"global"`
	Measurement []measurement `mapstructure:"measurement"`
}

type RawDataClient struct {
	Config   *Config
	Scope    *logUtil.Scope
	Queue    *queue.Queue
	DataPipeClient *datapipe.DataPipeClient
}

func NewRawDataClient() *RawDataClient {
	return &RawDataClient{}
}

func ReadConfig(filename string) (*Config, error) {
	var agentConfig Config
	viper.SetEnvPrefix("Test")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&agentConfig)
	if err != nil {
		panic(errors.New("Unmarshal configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(agentConfig, "", "  "); err == nil {
			// logger.Debug(fmt.Sprintf("Transmitter configuration: %s", string(transmitterConfBin)))
			fmt.Println(string(transmitterConfBin))
		}
	}
	return &agentConfig, nil
}

type inputLib struct {
}

var gDClient *RawDataClient

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gDClient.Queue = agentQueue
	gDClient.DataPipeClient.Queue = agentQueue
}

func (i inputLib) getNodes() []string {
	var nodesName []string
	resp, err := gDClient.DataPipeClient.GetNodes()
	if err != nil {
		return []string{}
	}
	if resp.Status.Code != 0 {
		gDClient.Scope.Errorf(fmt.Sprintf("failed to get nodes: %s", resp.Status.GetMessage()))
		return []string{}
	}
	for _, n := range resp.Nodes {
		nodesName = append(nodesName, n.Name)
	}
	return nodesName
}

func (i inputLib) Gather() error {
	st, _ := ptypes.TimestampProto(time.Now().Local())
	tr := utils.GetTimeRange(st, nil, 15, false, 60)
	// promQL := prometheus.NewPromQL("(sum(up{job=~\".*apiserver.*\"} == 1) / count(up{job=~\".*apiserver.*\"}))", "query_range", nil, gDClient.DataPipeClient)
	nodesName := i.getNodes()

	for _, m := range gDClient.Config.Measurement {
		var iFields []*InfluxConvert.InfluxField
		var exprs []string
		if strings.Contains(m.Expr, "\"$node\"") {
			for _, nName := range nodesName {
				expr := strings.Replace(m.Expr, "\"$node\"", fmt.Sprintf("\"%s\"", nName), -1)
				exprs = append(exprs, expr)
				gDClient.Scope.Infof(fmt.Sprintf("measurement: %v", expr))
			}
		} else {
			exprs = append(exprs, m.Expr)
		}
		for _, e := range exprs {
			promQL := prometheus.NewPromQL(e, "query_range", nil, gDClient.DataPipeClient)
			promQLResp, err := promQL.GetRawData(tr, 1)
			if err != nil{
				gDClient.Scope.Errorf(fmt.Sprintf("failed to get rawData %s %s, %v", m.Name, e, err))
				continue
			}
			if promQLResp != nil && promQLResp.Status.Code != 0 {
				gDClient.Scope.Errorf(fmt.Sprintf("failed to get rawData %s %s, %s", m.Name, e, promQLResp.Status.GetMessage()))
				continue
			}
			for k, v := range m.Element{
				iField := &InfluxConvert.InfluxField{Name: k, Type: InfluxConvert.StringToDataType(v.ElementType)}
				iFields = append(iFields, iField)
			}
			influxCV := InfluxConvert.NewInflux(gDClient.Config.Global.TargetDatabase, m.Name, m.Tags, iFields, promQLResp)
			rawWriteData, err := influxCV.GetWriteRequest()
			if err != nil{
				gDClient.Scope.Errorf(fmt.Sprintf("failed to covert data for write raw data %s, %v", m.Name, err))
				continue
			}
			data, err := json.Marshal(rawWriteData)
			gDClient.Scope.Debugf(fmt.Sprintf("measurement: %s, rawWriteData: %s", m.Name, string(data)))
			err = gDClient.DataPipeClient.WriteRawData(rawWriteData)
			if err != nil{
				gDClient.Scope.Errorf(fmt.Sprintf("failed to write raw data %s, %v", m.Name, err))
				continue
			}
			gDClient.Scope.Infof("succeed to write raw data %s", m.Name)
		}
	}
	// gDClient.Scope.Debugf(fmt.Sprintf("promQLResp: %v", promQLResp))
	/*
	bData, _ := json.Marshal(promQLResp)
	err = ioutil.WriteFile("/tmp/kubeStatus", bData, 0644)
	if err != nil {
		fmt.Println(err)
	}
	*/
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
	gDClient = NewRawDataClient()
	gDClient.Scope = scope
	gDClient.DataPipeClient = datapipe.NewDataPipeClient()

	viper.SetEnvPrefix("InputLibDataPipePrometheus")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	agentConfig, err := ReadConfig(config)
	if err != nil {
		scope.Errorf(fmt.Sprintf("Failed to load configuration, %v", err))
		return err
	}
	gDClient.Config = agentConfig
	gDClient.DataPipeClient.SetDataPipeAddress(agentConfig.Global.TargetAddress)
	gDClient.DataPipeClient.Scope = scope
	return nil
}

var InputLib inputLib
