package fedemeter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"crypto/tls"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"time"
	"context"
	"math"
)

type jeriParameter struct {
	tsfrom          int64
	tsto            int64
	granularity     int64
	fill_days       int
	ri_calculator   bool
}

type costParameter struct {
	tsfrom          int64
	tsto            int64
	granularity     int64
}

func NewFedermeter(apiUrl string, username string, password string, logger *logUtil.Scope) *Fedemeter {
	return &Fedemeter{apiUrl: apiUrl, user: username, password: password, logger: logger}
}

func (f *Fedemeter) GetApiInfo() (*map[string]string, error) {
	res, err := f.request("GET", fmt.Sprintf("%s/", f.apiUrl), nil, nil)
	if err != nil {
		f.logger.Errorf("Failed to get api server info, %v", err)
		return nil, err
	}
	apiServer := make(map[string]string)
	err = json.Unmarshal(res, &apiServer)
	return &apiServer, nil
}

func (f *Fedemeter) request(method string, url string, requestBody []byte, parameters interface{}) ([]byte, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	if len(f.user) > 0 {
		request.SetBasicAuth(f.user, f.password)
	}
	if err != nil {
		f.logger.Errorf("Failed to create request: %v", err)
		return nil, err
	}
	defer request.Body.Close()
	// Parameters
	if parameters != nil {
		q := request.URL.Query()
		refK := reflect.TypeOf(parameters).Elem()
		refV := reflect.ValueOf(parameters).Elem()
		for i := 0; i < refK.NumField(); i++ {
			value := ""
			if refV.Field(i).IsValid() == false  {
				continue
			}
			switch refV.Field(i).Kind() {
			case reflect.Int64, reflect.Int:
				value = fmt.Sprintf("%d", refV.Field(i).Int())
			case reflect.Bool:
				if refV.Field(i).Bool() == true {
					value = "true"
				} else {
					value = "false"
				}
			default:
				value = refV.Field(i).String()
			}
			q.Add(refK.Field(i).Name, value)
		}
		request.URL.RawQuery = q.Encode()
		f.logger.Infof("request: %s", request.URL.String())
	}
	client := &http.Client{Timeout: 120 * time.Second, Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		f.logger.Errorf("Failed to send http request with error %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("%d %s", resp.StatusCode, resp.Status)
	}

	stream, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.logger.Errorf("Failed to read response result, %v", err)
		return nil, err
	}
	return stream, nil
}

func (f *Fedemeter) ListProviders() (*fedProviderList, error) {
	var fedProviderList fedProviderList
	res, err := f.request("GET", fmt.Sprintf("%s/list/providers", f.apiUrl), nil, nil)
	if err != nil {
		f.logger.Errorf("Failed to list providers, %v", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &fedProviderList)
	return &fedProviderList, nil
}

func (f *Fedemeter) ListRegions() (*fedRegionList, error) {
	var fedRegion fedRegionList
	res, err := f.request("GET", fmt.Sprintf("%s/list/regions", f.apiUrl), nil, nil)
	if err != nil {
		f.logger.Errorf("Failed to list regions, %v", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &fedRegion)
	return &fedRegion, nil
}

func (f *Fedemeter) ListInstances(provider string) (*fedInstanceList, error) {
	var fedInstances fedInstanceList
	res, err := f.request("GET", fmt.Sprintf("%s/list/%s/instances", f.apiUrl, provider), nil, nil)
	if err != nil {
		f.logger.Errorf("Failed to list instance, %v", err)
		return nil, err
	}

	err = json.Unmarshal(res, &fedInstances)
	return &fedInstances, nil
}

func (f *Fedemeter) GetRecommenderationJri(fedJriRequest *FedRecommendationJri) (*FedRecommendationJriResp, error) {
	var fedJriResp FedRecommendationJriResp
	reqBody, err := json.Marshal(fedJriRequest)
	if err != nil {
		return nil, err
	}
	res, err := f.request("GET", fmt.Sprintf("%s/recommendations/jri", f.apiUrl), reqBody, nil)
	if err != nil {
		f.logger.Errorf("Failed to get recommendations jre, %v", err)
		return nil, err
	}

	err = json.Unmarshal(res, &fedJriResp)
	return &fedJriResp, nil
}

func (f *Fedemeter) GetCostHistorical(tsFrom int64, tsTo int64, granularity int64, fedCostReq *FedCostMetricReq) (*FedCostMetricResp, error) {
	var fedCostResp FedCostMetricResp
	reqBody, err := json.Marshal(fedCostReq)
	if err != nil {
		return nil, err
	}
	parms := &costParameter{tsfrom: tsFrom, tsto: tsTo, granularity: granularity}
	res, err := f.request("PUT", fmt.Sprintf("%s/resources/historical/cost/", f.apiUrl), reqBody, parms)
	if err != nil {
		f.logger.Errorf("Failed to get historical cost, %v", err)
		return nil, err
	}
	f.logger.Debugf("GetCostHistorical: %s", utils.InterfaceToString(string(res)))
	err = json.Unmarshal(res, &fedCostResp)
	if err != nil {
		return nil, err
	}
	for _, n := range fedCostResp.Cluster.Provider.Namespace {
		for i, c := range n.Costs {
			n.Costs[i].Timestampe = int64(math.Floor(float64(c.Timestampe) / float64(granularity)))
		}
		for x, a := range n.App {
			for i, c := range a.Costs {
				n.App[x].Costs[i].Timestampe = int64(math.Floor(float64(c.Timestampe) / float64(granularity)))
			}
		}
	}
	f.logger.Debugf("CostHistorical: %s", utils.InterfaceToString(fedCostResp))
	return &fedCostResp, nil
}

func (f *Fedemeter) GetCostPredicted(tsFrom int64, tsTo int64, granularity int64, fedCostReq *FedCostMetricReq) (*FedCostMetricResp, error) {
	var fedCostResp FedCostMetricResp
	reqBody, err := json.Marshal(fedCostReq)
	if err != nil {
		return nil, err
	}
	parms := &costParameter{tsfrom: tsFrom, tsto: tsTo, granularity: granularity}
	res, err := f.request("PUT", fmt.Sprintf("%s/resources/predictions/cost/", f.apiUrl), reqBody, parms)
	if err != nil {
		f.logger.Errorf("Failed to get predicted cost, %v", err)
		return nil, err
	}
	f.logger.Debugf("GetCostPredicted: %s", utils.InterfaceToString(string(res)))
	err = json.Unmarshal(res, &fedCostResp)
	if err != nil {
		return nil, err
	}
	for _, n := range fedCostResp.Cluster.Provider.Namespace {
		for i, c := range n.Costs {
			n.Costs[i].Timestampe = int64(math.Floor(float64(c.Timestampe) / float64(granularity)))
		}
		for x, a := range n.App {
			for i, c := range a.Costs {
				n.App[x].Costs[i].Timestampe = int64(math.Floor(float64(c.Timestampe) / float64(granularity)))
			}
		}
	}
	f.logger.Debugf("CostPredicted: %s", utils.InterfaceToString(fedCostResp))
	return &fedCostResp, nil
}

func (f *Fedemeter) GetRecommenderationJeri(tsFrom int64, tsTo int64, granularity int64, fillDays int, fedJeriRequest *FedRecommendationJeri, jeri bool) (*FedRecommendationJeriResp, error) {
	var fedJriResp FedRecommendationJeriResp
	reqBody, err := json.Marshal(fedJeriRequest)
	if err != nil {
		return nil, err
	}
	parms := &jeriParameter{tsfrom: tsFrom, tsto: tsTo, granularity: granularity, fill_days: fillDays, ri_calculator: jeri}
	res, err := f.request("PUT", fmt.Sprintf("%s/recommendations/jeri", f.apiUrl), reqBody, parms)
	if err != nil {
		f.logger.Errorf("Failed to get recommendations jeri, %v", err)
		return nil, err
	}

	err = json.Unmarshal(res, &fedJriResp)
	if err != nil {
		return nil, err
	}
	return &fedJriResp, nil
}

func (f *Fedemeter) Calculate (fedProviders *FedProviders) (*FedCalculatorResp, error) {
	var fedCalculatorResp FedCalculatorResp
	reqBody, err := json.Marshal(fedProviders)
	if err != nil {
		return nil, err
	}
	res, err := f.request("PUT", fmt.Sprintf("%s/calculators/", f.apiUrl), reqBody, nil)
	if err != nil {
		f.logger.Errorf("Failed to calculate node cost with error %v", err)
		return nil, err
	}

	err = json.Unmarshal(res, &fedCalculatorResp)
	if err != nil {
		return nil, err
	}
	return &fedCalculatorResp, nil
}
