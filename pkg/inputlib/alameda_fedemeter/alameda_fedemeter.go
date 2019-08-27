package main

import (
	"time"
	"strconv"
	"errors"
	"encoding/json"
	"fmt"
	"strings"
	DP "github.com/containers-ai/federatorai-agent/pkg/datahub"
	FedAPI "github.com/containers-ai/federatorai-agent/pkg/fedemeter"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/alameda/operator/datahub"
	"github.com/sheerun/queue"
	"github.com/spf13/viper"
	"github.com/containers-ai/federatorai-agent/pkg/adapter"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	FedRaw "github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_fedemeter/influx"
)

type CostAnalysisConf struct {
	DataHub  *datahub.Config  `mapstructure:"datahub"`
	Fedemeter * struct {
		Url string `mapstructure:"url"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"fedemeter"`
	CostAnalysis * struct {
		CalculateCurrent bool `mapstructure:"calculate_current"`
		CalculateCurrentUnit string `mapstructure:"calculate_current_unit"`
	} `mapstructure:"cost_analysis"`
	Recommendation * struct {
		Ri bool `mapstructure:"ri"`
		Granularity string `mapstructure:"granularity"`
		FillDays string `mapstructure:"fill_days"`
	} `mapstructure:"recommendation"`
}

var gFedermeter *CostAnalyst
var logger *logUtil.Scope

type inputLib struct {
}

func getTimeFromTo(starttime *timestamp.Timestamp, durationDays int) (int64, int64) {
	var ts *timestamp.Timestamp
	if starttime == nil {
		ts = ptypes.TimestampNow()
	} else {
		ts = starttime
	}
	sTm, _ := ptypes.Timestamp(ts)
	du, _ := time.ParseDuration(fmt.Sprintf("%dh", durationDays * 24))
	sTm = sTm.Add(du)
	es, _ := ptypes.TimestampProto(sTm)
	return ts.Seconds, es.Seconds
}

func (i inputLib) Gather() error {
	fill_days, _ := strconv.Atoi(gFedermeter.Conf.Recommendation.FillDays)
	st, es := getTimeFromTo(nil, fill_days)
	granularity, _ := strconv.ParseInt(gFedermeter.Conf.Recommendation.Granularity, 10, 64)
	logger.Debugf("fill days: %d, st: %d, es: %d, granularity: %d(%d) object: %p", fill_days, st, es, granularity, es-st, gFedermeter)

	// Get nodes
	tm := ptypes.TimestampNow()
	nodes, err := gFedermeter.DPClient.GetNodes()
	if err != nil {
		logger.Errorf("Failed to get nodes with error %v", err)
	}
	adpFed := adapterK8sToFedmeter.NewAdapterNodes(nodes.Nodes)
	if gFedermeter.Conf.CostAnalysis.CalculateCurrent == true {
		fedCal, err := adpFed.GenerateFedemeterCalculates(gFedermeter.Conf.CostAnalysis.CalculateCurrentUnit)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to generate fedemeter calculation request format %v", err))
			return err
		}
		logger.Debugf("calculation format: %s", utils.InterfaceToString(fedCal))
		// Get cluster calculation result
		fedCalResp, err := gFedermeter.FedApi.Calculate(fedCal)
		if err != nil {
			logger.Errorf(fmt.Sprintf("Failed to get fedemeter calculation response %v", err))
			return err
		}
		logger.Debugf("Calculation response: %s", utils.InterfaceToString(fedCalResp))

		iMeasurement := FedRaw.NewInfluxMeasurement("alameda_fedemeter", 0, nil, nil, fedCalResp, 3600, false)
		fedCalInstanceRawData, err := iMeasurement.GetWriteRequest(&timestamp.Timestamp{Seconds: tm.Seconds, Nanos: tm.Nanos})
		if err != nil {
			logger.Errorf(fmt.Sprintf("Failed to generate fedemeter calculation instances %v", err))
			return err
		}
		err = gFedermeter.DPClient.WriteRawData(fedCalInstanceRawData)
		if err != nil {
			logger.Errorf( fmt.Sprintf("Failed to write fedemeter calculation instances %v", err))
			return err
		}
		logger.Info("Succeed to write fedemeter calculation instances")

		iStorageMeasurement := FedRaw.NewInfluxMeasurement("alameda_fedemeter", 1, nil, nil, fedCalResp, 3600, false)
		fedCalStorageRawData, err := iStorageMeasurement.GetWriteRequest(&timestamp.Timestamp{Seconds: tm.Seconds, Nanos: tm.Nanos})
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to generate fedemeter calculation storage %v", err))
			return err
		}
		err = gFedermeter.DPClient.WriteRawData(fedCalStorageRawData)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to write fedemeter calculation storage %v", err))
			return err
		}
		logger.Info("Succeed to write fedemeter calculation storage")
	}

	for i := 0; i < 2; i++ {
		var enableRi bool
		if i == 0 {
			enableRi = true
		} else {
			enableRi = false
		}
		fedRC, err := adpFed.GenerateFedemeterRecommendationNodes("jeri", "hour")
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to generate fedemeter recommendation request format %v", err))
			return err
		}

		fedRCResp, err := gFedermeter.FedApi.GetRecommenderationJeri(st, es, granularity, fill_days, fedRC, enableRi)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to get fedemeter recommendation %v", err))
			return err
		}
		logger.Debugf(fmt.Sprintf("Recommendation: %s", utils.InterfaceToString(fedRCResp)))

		iRecommendationMeasurement := FedRaw.NewInfluxMeasurement("alameda_fedemeter", 2, nil, nil, fedRCResp, granularity, enableRi)
		fedCalRecommendationRawData, err := iRecommendationMeasurement.GetWriteRequest(&timestamp.Timestamp{Seconds: tm.Seconds, Nanos: tm.Nanos})
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to generate fedemeter recommendation %v", err))
			return err
		}
		logger.Debugf("Recommendation raw data: %s", utils.InterfaceToString(fedCalRecommendationRawData))

		err = gFedermeter.DPClient.WriteRawData(fedCalRecommendationRawData)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to write fedemeter recommendation %v", err))
			return err
		}
		logger.Info(fmt.Sprintf("Succeed to write recommendation fill_dys: %d, start time: %d, granularity: %d, ri calculate: %v", fill_days, st, granularity, enableRi))
	}
	return nil
}

func (i inputLib) LoadConfig(config *string, scope *logUtil.Scope) error {
	gFedermeter = NewCostAnalyst(config, scope)
	logger = scope
	return nil
}

func (i inputLib) SetAgentQueue(agentQueue *queue.Queue) {
	gFedermeter.Queue = agentQueue
}

func (i inputLib) Close() {
	gFedermeter.Scope.Debugf("Free cost analyst")
	gFedermeter = &CostAnalyst{}
}

type CostAnalyst struct {
	Conf *CostAnalysisConf
	Queue *queue.Queue
	Scope *logUtil.Scope
	DPClient *DP.DataHubClient
	FedApi *FedAPI.Fedemeter
}

func NewCostAnalyst(confPath *string, logger *logUtil.Scope) *CostAnalyst {
	var fedConf CostAnalysisConf
	viper.SetEnvPrefix("InputLibFedemeterCostAnalysis")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	logger.Debugf("Load config path:(%p) %s", confPath, *confPath)

	viper.SetConfigFile(*confPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library fedemeter cost analysis configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&fedConf)
	if err != nil {
		panic(errors.New("Unmarshal input library fedemeter cost analysis configuration failed: " + err.Error()))
	} else {
		if _, err := json.MarshalIndent(fedConf, "", "  "); err == nil {
			// scope.Debugf(fmt.Sprintf("Input library fedemeter cost analysis configuration: %s", string(transmitterConfBin)))
		} else {
			fmt.Printf("failed to display fedemeter cost analysis configuration")
		}
	}

	dp := DP.NewDataHubClient()
	dp.Scope = logger
	dp.DataHub.DataHub = fedConf.DataHub
	fed := FedAPI.NewFedermeter(fedConf.Fedemeter.Url, fedConf.Fedemeter.Username, fedConf.Fedemeter.Password, logger)
	return &CostAnalyst{Conf: &fedConf, Scope: logger, DPClient: dp, FedApi: fed}
}

func (c *CostAnalyst) SetQueue(queue *queue.Queue) {
	c.Queue = queue
}

var InputLib inputLib
