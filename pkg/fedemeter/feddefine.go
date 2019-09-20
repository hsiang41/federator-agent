package fedemeter

import (
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"encoding/json"
)

type Fedemeter struct {
	apiUrl  string
	user    string
	password string
	logger  *logUtil.Scope
}

type FedInstances struct {
	Instancetype    string  `json:"instancetype"`
	Operatingsystem string  `json:"operatingsystem"`
	Preinstalledsw  string  `json:"preinstalledsw"`
	Instancenum     string  `json:"instancenum"`
	Period          string  `json:"period"`
	Unit            string  `json:"unit"`
	Nodename        string  `json:"nodename"`
	Nodetype        string  `json:"nodetype"`
}

type FedGpu struct {
	Gpuinstance string
	Gpucores    int
	Gpunum      int
	Period      int
	Unit        string
}

type FedStorage struct {
	Volumetype  string      `json:"volumetype"`
	Storagesize string      `json:"storagesize"`
	Storagenum  string      `json:"storagenum"`
	Period      string      `json:"period"`
	Unit        string      `json:"unit"`
}

type FedProvider struct {
	Region      string        `json:"region"`
	Instances   *FedInstances `json:"instances"`
	Gpu         *FedGpu       `json:"gpu,omitempty"`
	Storage     []*FedStorage `json:"storage"`
}

type FedProviders struct {
	Calculator  []map[string][]*FedProvider `json:"calculator"`
}

type fedProvidersResp struct {
	Instances struct {
		Nodename        string  `json:"nodename"`
		Instancetype    string  `json:"instancetype"`
		Nodetype        string  `json:"nodetype"`
		Operatingsystem string  `json:"operatingsystem"`
		Preinstalledsw  string  `json:"preinstalledsw"`
		Instancenum     string  `json:"instancenum"`
		Period          string  `json:"period"`
		Unit            string  `json:"unit"`
		Description     string  `json:"description"`
		Cost            float64 `json:"cost"`
		CPU             float64 `json:"cpu"`
		Memory          float64 `json:"memory"`
		Displayname     string  `json:"displayname"`
	} `json:"instances"`
	Storage []struct {
		Volumetype  string  `json:"volumetype"`
		Storagesize string  `json:"storagesize"`
		Storagenum  string  `json:"storagenum"`
		Period      string  `json:"period"`
		Unit        string  `json:"unit"`
		Description string  `json:"description"`
		Cost        float64 `json:"cost"`
		Displayname string  `json:"displayname"`
	} `json:"storage"`
	Totalcost float64 `json:"totalcost"`
	Region    string  `json:"region"`
	Status    string  `json:"status"`
	Gpu       struct {
		Gpuinstance string  `json:"gpuinstance"`
		Gpucores    string  `json:"gpucores"`
		Gpunum      string  `json:"gpunum"`
		Period      string  `json:"period"`
		Unit        string  `json:"unit"`
		Description string  `json:"description"`
		Cost        float64 `json:"cost"`
		Displayname string  `json:"displayname"`
	} `json:"gpu,omitempty"`
}

type FedCalculatorResp struct {
	Calculator []map[string][] *fedProvidersResp  `json:"calculator"`
	Count      int  `json:"count"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	Offset     int  `json:"offset"`
	Total      int  `json:"total"`
	Incomplete bool `json:"incomplete"`
}

type fedProviderList struct {
	Providers []string `json:"providers"`
	Count     int      `json:"count"`
	Limit     int      `json:"limit"`
	Page      int      `json:"page"`
	Offset    int      `json:"offset"`
	Total     int      `json:"total"`
}

type fedRegionList struct {
	Regions []map[string]string `json:"regions"`
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type fedInstanceList struct {
	Instances []struct {
		Aws string `json:"aws"`
	} `json:"instances"`
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type FedRecommendation struct {
	Acceptance []map[string]string
	Configuration map[string][]*FedProvider
}

type FedRecommendationJri struct {
	Recommender []*FedRecommendation
}

type FedRecommendationJeriResource struct {
	Acceptance []map[string]string      `json:"acceptance"`
	Clustername string                  `json:"clustername"`
	Type        string                  `json:"type"`
	Nodesinfo map[string][]*FedProvider `json:"nodesinfo"`
	Category    string                  `json:"category"`
}

type FedRecommendationJeri struct {
	Resource []*FedRecommendationJeriResource `json:"resource"`
}

type FedRecommendationJriResp struct {
	Recommender []struct {
		Aws []struct {
			Instances struct {
				Nodename        string  `json:"nodename"`
				Instancetype    string  `json:"instancetype"`
				Nodetype        string  `json:"nodetype"`
				Operatingsystem string  `json:"operatingsystem"`
				Preinstalledsw  string  `json:"preinstalledsw"`
				Instancenum     string  `json:"instancenum"`
				Period          int     `json:"period"`
				Unit            string  `json:"unit"`
				Sourcetype      string  `json:"sourcetype"`
				Description     string  `json:"description"`
				Cost            float64 `json:"cost"`
				Displayname     string  `json:"displayname"`
			} `json:"instances"`
			Storage []struct {
				Volumetype  string  `json:"volumetype"`
				Storagesize string  `json:"storagesize"`
				Storagenum  string  `json:"storagenum"`
				Period      string  `json:"period"`
				Unit        string  `json:"unit"`
				Description string  `json:"description"`
				Cost        float64 `json:"cost"`
				Displayname string  `json:"displayname"`
			} `json:"storage"`
			Totalcost float64 `json:"totalcost"`
			Region    string  `json:"region"`
			Status    string  `json:"status"`
		} `json:"aws,omitempty"`
	} `json:"recommender"`
	Count      int  `json:"count"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	Offset     int  `json:"offset"`
	Total      int  `json:"total"`
	Incomplete bool `json:"incomplete"`
}

type FedJeriInstance struct {
	MasterNum         int     `json:"master_num"`
	WorkerNum         int     `json:"worker_num"`
	MasterStorageSize float64 `json:"master_storage_size"`
	WorkerStorageSize float64 `json:"worker_storage_size"`
	StorageCost       float64 `json:"storage_cost"`
	AccCost           float64 `json:"acc_cost"`
	Displayname       string  `json:"displayname"`
	Timestamp         int     `json:"timestamp"`
	OndemandCost      float64 `json:"ondemand_cost"`
	OndemandNum       int     `json:"ondemand_num"`
	RiCost            float64 `json:"ri_cost"`
	RiNum             int     `json:"ri_num"`
	MasterRiNum       int     `json:"master_ri_num"`
	WorkerRiNum       int     `json:"worker_ri_num"`
	MasterOndemandNum int     `json:"master_ondemand_num"`
	WorkerOndemandNum int     `json:"worker_ondemand_num"`
}

type fedJeriProvider struct {
	Region   string                        `json:"region"`
	Profider map[string][]*FedJeriInstance `json:""`
}

type FedRecommendationJeriResp struct {
	Resource map[string] map[string][]map[string]map[string] json.RawMessage `json:"resource"`
	Count      int  `json:"count"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	Offset     int  `json:"offset"`
	Total      int  `json:"total"`
	Incomplete bool `json:"incomplete"`
}

type FedCostMetricResp struct {
	Cluster struct {
		Clustername string `json:"clustername"`
		Providers   [] struct {
			Providername string `json:"providername"`
			Namespace    []struct {
				Namespacename string `json:"namespacename"`
				Costs         []struct {
					Workloadcost   string `json:"workloadcost"`
					Costpercentage string `json:"costpercentage"`
					Timestampe     int64  `json:"timestampe"`
				} `json:"costs"`
				Apps []struct {
					Appname string `json:"appname"`
					Costs   []struct {
						Workloadcost   string `json:"workloadcost"`
						Costpercentage string `json:"costpercentage"`
						Timestampe     int64  `json:"timestampe"`
					} `json:"costs"`
				} `json:"apps"`
			} `json:"namespace"`
		} `json:"providers"`
	} `json:"cluster"`
	Count      int  `json:"count"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	Offset     int  `json:"offset"`
	Total      int  `json:"total"`
	Incomplete bool `json:"incomplete"`
}

type FedCostMetricResource struct {
	Category    string `json:"category"`
	Type        string `json:"type"`
	Clustername string `json:"clustername"`
	Nodesinfo   map[string][]*FedProvider `json:"nodesinfo"`
}

type FedCostMetricReq struct {
	Resource [] *FedCostMetricResource `json:"resource"`
}