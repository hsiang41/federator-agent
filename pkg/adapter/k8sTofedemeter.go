package adapterK8sToFedmeter

import (
	"fmt"
	"strings"
	datahubV1a1pha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	Fed "github.com/containers-ai/federatorai-agent/pkg/fedemeter"
)

var AwsRegionMap = map[string]string {
	"ap-south-1": "Asia Pacific (Mumbai)",
	"ap-northeast-1": "Asia Pacific (Tokyo)",
	"ap-northeast-2": "Asia Pacific (Seoul)",
	"ap-northeast-3": "Asia Pacific (Osaka-Local)",
	"ap-southeast-1": "Asia Pacific (Singapore)",
	"ap-southeast-2": "Asia Pacific (Sydney)",
	"aws-govcloud-1": "AWS GovCloud (US)",
	"aws-govcloud-2": "AWS GovCloud (US-EAST)",
	"ca-central-1": "Canada (Central)",
	"eu-central-1": "EU (Frankfurt)",
	"eu-west-1": "EU (Ireland)",
	"eu-west-2": "EU (London)",
	"eu-west-3": "EU (Paris)",
	"eu-west-4": "EU (Stockholm)",
	"sa-east-1": "South America (Sao Paulo)",
	"me-south-1": "Middle East (Bahrain)",
	"us-east-1": "US East (N. Virginia)",
	"us-east-2": "US East (Ohio)",
	"us-west-1": "US West (N. California)",
	"us-west-2": "US West (Oregon)",
}

var GcpRegionMap = map[string]string {
	"asia-east1-a":	"Asia East 1a",
	"asia-east1-b": "Asia East 1b",
	"asia-east1-c": "Asia East 1c",
	"asia-east2-a":	"Asia East 2a",
	"asia-east2-b":	"Asia East 2b",
	"asia-east2-c":	"Asia East 2c",
	"asia-northeast1-a":	"Asia Northeast 1a",
	"asia-northeast1-b":	"Asia Northeast 1b",
	"asia-northeast1-c":	"Asia Northeast 1c",
	"asia-south1-a":	"Asia South 1a",
	"asia-south1-b":	"Asia South 1b",
	"asia-south1-c":	"Asia South 1c",
	"asia-southeast1-a":	"Asia Southeast 1a",
	"asia-southeast1-b":	"Asia Southeast 1b",
	"asia-southeast1-c":	"Asia Southeast 1c",
	"australia-southeast1-a":	"Australia Southeast 1a",
	"australia-southeast1-b":	"Australia Southeast 1b",
	"australia-southeast1-c":	"Australia Southeast 1c",
	"europe-north1-a":	"Europe North 1a",
	"europe-north1-b":	"Europe North 1b",
	"europe-north1-c":	"Europe North 1c",
	"europe-west1-b":	"Europe West 1b",
	"europe-west1-c":	"Europe West 1c",
	"europe-west1-d":	"Europe West 1d",
	"europe-west2-a":	"Europe West 2a",
	"europe-west2-b":	"Europe West 2b",
	"europe-west2-c":	"Europe West 2c",
	"europe-west3-a":	"Europe West 3a",
	"europe-west3-b":	"Europe West 3b",
	"europe-west3-c":	"Europe West 3c",
	"europe-west4-a":	"Europe West 4a",
	"europe-west4-b":	"Europe West 4b",
	"europe-west4-c":	"Europe West 4c",
	"northamerica-northeast1-a":	"Northamerica Northeast 1a",
	"northamerica-northeast1-b":	"Northamerica Northeast 1b",
	"northamerica-northeast1-c":	"Northamerica Northeast 1c",
	"southamerica-east1-a":	"Southamerica East 1a",
	"southamerica-east1-b":	"Southamerica East 1b",
	"southamerica-east1-c":	"Southamerica East 1c",
	"us-central1-a":	"US Central 1a",
	"us-central1-b":	"US Central 1b",
	"us-central1-c":	"US Central 1c",
	"us-central1-f":	"US Central 1f",
	"us-east1-b":	"US East 1b",
	"us-east1-c":	"US East 1c",
	"us-east1-d":	"US East 1d",
	"us-east4-a":	"US East 4a",
	"us-east4-b":	"US East 4b",
	"us-east4-c":	"US East 4c",
	"us-west1-a":	"US West 1a",
	"us-west1-b":	"US West 1b",
	"us-west1-c":	"US West 1c",
	"us-west2-a":	"US West 2a",
	"us-west2-b":	"US West 2b",
	"us-west2-c":	"US West 2c",
}


type AdapterNodes struct {
	Nodes []*datahubV1a1pha1.Node
}

func NewAdapterNodes(nodes []*datahubV1a1pha1.Node) *AdapterNodes {
	return &AdapterNodes{Nodes: nodes}
}

func covertRegion(provider string, region string) string {
	var ok bool
	var toRegion string
	switch strings.ToLower(provider) {
	case "aws":
		toRegion, ok = AwsRegionMap[region]
		if ok == false {
			toRegion = region
		}
	case "gcp":
		toRegion, ok = GcpRegionMap[region]
		if ok == false {
			toRegion = region
		}
	}
	return toRegion
}

func covertNameToFedNodeName(nodename string) string {
	name := strings.Split(nodename, ".")[0]
	if len(name) >= len("ip-") && name[0:3] == "ip-" {
		return name[3:]
	} else {
		return name
	}
}

func getVolumeType(provider string) string {
	switch strings.ToLower(provider) {
	case "aws":
		return "General Purpose"
	case "gce":
		return "CP-COMPUTEENGINE-STORAGE-PD-SSD"
	case "azure":
		return "standardssd-e6"
	}
	return provider
}

func getProviderRaw(provider string) string {
	if strings.ToLower(provider) == "gce" {
		return "gcp"
	}
	return provider
}

func getProviderOS(provider string, os string) string {
	if strings.ToLower(provider) == "gce" {
		return "free"
	}
	return os
}

func (a *AdapterNodes) GenerateFedemeterCalculates(unit string) (*Fed.FedProviders, error) {
	var fedProviders Fed.FedProviders
	fedP := make(map[string][]*Fed.FedProvider, 0)
	for _, n := range a.Nodes {
		var fedProvider Fed.FedProvider
		var fedStorage Fed.FedStorage
		fedProvider.Region = covertRegion(n.Provider.Provider, n.Provider.Region)
		fedProvider.Instances = &Fed.FedInstances{}
		fedProvider.Instances.Period = "1"
		fedProvider.Instances.Unit = unit
		fedProvider.Instances.Instancenum = "1"
		fedProvider.Instances.Nodename = covertNameToFedNodeName(n.Name)
		fedProvider.Instances.Instancetype = n.Provider.InstanceType
		fedProvider.Instances.Nodetype = n.Provider.Role
		fedProvider.Instances.Operatingsystem = strings.Title(n.Provider.Os)
		fedProvider.Instances.Preinstalledsw = "NA"
		fedStorage.Volumetype = getVolumeType(n.Provider.Provider)
		fedStorage.Storagenum = "1"
		fedStorage.Period = "1"
		fedStorage.Unit = unit
		fedStorage.Storagesize = fmt.Sprintf("%d", n.Provider.StorageSize / (1000 * 1000 * 1000))
		fedProvider.Storage = append(fedProvider.Storage, &fedStorage)
		fedP[getProviderRaw(n.Provider.Provider)] = append(fedP[getProviderRaw(n.Provider.Provider)], &fedProvider)
	}
	fedProviders.Calculator = append(fedProviders.Calculator, fedP)
	return &fedProviders, nil
}

func (a *AdapterNodes) GenerateFedemeterRecommendationNodes (jeriType string, unit string) (*Fed.FedRecommendationJeri, error) {
	var fedRecJeri Fed.FedRecommendationJeri
	var fedRec     Fed.FedRecommendationJeriResource
	fedNodes, err := a.GenerateFedemeterCalculates(unit)
	if err != nil {
		return nil, err
	}
	fedRec.Category = "predicted"
	fedRec.Type = jeriType
	fedRec.Clustername = "Royal Grass"
	fedRec.Nodesinfo = fedNodes.Calculator[0]
	fedRecJeri.Resource = append(fedRecJeri.Resource, &fedRec)
	return &fedRecJeri, nil
}

func (a *AdapterNodes) GenerateFedemeterCostRequest(clusterName string, category string, costType string, unit string)(*Fed.FedCostMetricReq, error) {
	var fedCostReq Fed.FedCostMetricReq
	var fedCostReqResource Fed.FedCostMetricResource
	fedNodes, err := a.GenerateFedemeterCalculates(unit)
	if err != nil {
		return nil, err
	}
	fedCostReqResource.Nodesinfo = fedNodes.Calculator[0]
	fedCostReqResource.Clustername = clusterName
	fedCostReqResource.Category = category
	fedCostReqResource.Type = costType

	fedCostReq.Resource = append(fedCostReq.Resource, &fedCostReqResource)
	return &fedCostReq, nil
}