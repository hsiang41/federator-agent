package fedemeter

import (
	"fmt"
	"testing"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
)

var logger = logUtil.RegisterScope("fedemeter_test", "fedemeter testing", 0)
var fed *Fedemeter

func init() {
	// 54.218.143.157, 34.221.21.224
	fed = NewFedermeter("http://34.223.245.164:31000/fedemeter-api/v1", "fedemeter", "$6$pOwGiawPSjz7qLaN$fnMXEhwzWnUw.bOKohdAhB5K5iCCOJJaZXxQkhzH4URsHP8qLTT4QeBPUKjlOAeAHbKsqlf.fyuL2pNRmR6oQD1", logger)
}

func TestFedermeter_GetApiInfo(t *testing.T) {
	apiServerInfo, err := fed.GetApiInfo()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(apiServerInfo)
}

func TestFedermeter_ListProviders(t *testing.T) {
	fedProvider, err := fed.ListProviders()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.InterfaceToString(fedProvider))
}

func TestFedermeter_ListRegion(t *testing.T) {
	fedRegion, err := fed.ListRegions()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.InterfaceToString(fedRegion))
}

func TestFedermeter_GetRecommenderationJeri(t *testing.T) {
	fedRcJeri := &FedRecommendationJeri{}
	fedRcres := &FedRecommendationJeriResource{}
	fedRcres.Type = "jeri"
	fedRcres.Category = "predicted"
	fedRcres.Clustername = "Royal Grass"
	acceptance := make(map[string]string, 0)
	acceptance["provider"] = "*"
	fedRcres.Acceptance = append(fedRcres.Acceptance, acceptance)
	fedRcres.Nodesinfo = make(map[string][]*FedProvider)
	// Instance 1
	fedRcNode1 := &FedProvider{}
	fedRcNode1.Region = "US West (Oregon)"
	fedRcNode1.Instances = &FedInstances{}
	fedRcNode1.Instances.Nodename = "172-23-1-12"
	fedRcNode1.Instances.Instancetype = "m5.2xlarge"
	fedRcNode1.Instances.Nodetype = "master"
	fedRcNode1.Instances.Operatingsystem = "Linux"
	fedRcNode1.Instances.Preinstalledsw = "NA"
	fedRcNode1.Instances.Instancenum = "1"
	fedRcNode1.Instances.Period = "1"
	fedRcNode1.Instances.Unit = "hour"
	fedRcNode1Storage := &FedStorage{}
	fedRcNode1Storage.Unit = "hour"
	fedRcNode1Storage.Volumetype = "General Purpose"
	fedRcNode1Storage.Storagesize = "50"
	fedRcNode1Storage.Storagenum = "1"
	fedRcNode1Storage.Period = "1"
	fedRcNode1.Storage = append(fedRcNode1.Storage, fedRcNode1Storage)

	// Instance 2
	fedRcNode2 := &FedProvider{}
	fedRcNode2.Region = "US West (Oregon)"
	fedRcNode2.Instances = &FedInstances{}
	fedRcNode2.Instances.Nodename = "172-23-1-45"
	fedRcNode2.Instances.Instancetype = "m5.xlarge"
	fedRcNode2.Instances.Nodetype = "worker"
	fedRcNode2.Instances.Operatingsystem = "Linux"
	fedRcNode2.Instances.Preinstalledsw = "NA"
	fedRcNode2.Instances.Instancenum = "1"
	fedRcNode2.Instances.Period = "1"
	fedRcNode2.Instances.Unit = "hour"
	fedRcNode2Storage := &FedStorage{}
	fedRcNode2Storage.Unit = "hour"
	fedRcNode2Storage.Volumetype = "General Purpose"
	fedRcNode2Storage.Storagesize = "50"
	fedRcNode2Storage.Storagenum = "1"
	fedRcNode2Storage.Period = "1"
	fedRcNode2.Storage = append(fedRcNode2.Storage, fedRcNode2Storage)

	// Instance 3
	fedRcNode3 := &FedProvider{}
	fedRcNode3.Region = "US West (Oregon)"
	fedRcNode3.Instances = &FedInstances{}
	fedRcNode3.Instances.Nodename = "172-23-1-21"
	fedRcNode3.Instances.Instancetype = "m5.xlarge"
	fedRcNode3.Instances.Nodetype = "worker"
	fedRcNode3.Instances.Operatingsystem = "Linux"
	fedRcNode3.Instances.Preinstalledsw = "NA"
	fedRcNode3.Instances.Instancenum = "1"
	fedRcNode3.Instances.Period = "1"
	fedRcNode3.Instances.Unit = "hour"
	fedRcNode3Storage := &FedStorage{}
	fedRcNode3Storage.Unit = "hour"
	fedRcNode3Storage.Volumetype = "General Purpose"
	fedRcNode3Storage.Storagesize = "50"
	fedRcNode3Storage.Storagenum = "1"
	fedRcNode3Storage.Period = "1"
	fedRcNode3.Storage = append(fedRcNode3.Storage, fedRcNode3Storage)

	// Instance 4
	fedRcNode4 := &FedProvider{}
	fedRcNode4.Region = "US West (Oregon)"
	fedRcNode4.Instances = &FedInstances{}
	fedRcNode4.Instances.Nodename = "172-23-1-181"
	fedRcNode4.Instances.Instancetype = "m5.xlarge"
	fedRcNode4.Instances.Nodetype = "worker"
	fedRcNode4.Instances.Operatingsystem = "Linux"
	fedRcNode4.Instances.Preinstalledsw = "NA"
	fedRcNode4.Instances.Instancenum = "1"
	fedRcNode4.Instances.Period = "1"
	fedRcNode4.Instances.Unit = "hour"
	fedRcNode4Storage := &FedStorage{}
	fedRcNode4Storage.Unit = "hour"
	fedRcNode4Storage.Volumetype = "General Purpose"
	fedRcNode4Storage.Storagesize = "50"
	fedRcNode4Storage.Storagenum = "1"
	fedRcNode4Storage.Period = "1"
	fedRcNode4.Storage = append(fedRcNode4.Storage, fedRcNode4Storage)

	fedRcres.Nodesinfo["aws"] = append(fedRcres.Nodesinfo["aws"], fedRcNode1)
	fedRcres.Nodesinfo["aws"] = append(fedRcres.Nodesinfo["aws"], fedRcNode2)
	fedRcres.Nodesinfo["aws"] = append(fedRcres.Nodesinfo["aws"], fedRcNode3)
	fedRcres.Nodesinfo["aws"] = append(fedRcres.Nodesinfo["aws"], fedRcNode4)

	fedRcJeri.Resource = append(fedRcJeri.Resource, fedRcres)
	//fromTs := ptypes.TimestampNow()
	//toTs := &timestamp.Timestamp{Seconds: fromTs.Seconds + (7 * 24 * int64(time.Hour/time.Second))}

	fmt.Println(utils.InterfaceToString(fedRcJeri))
	fedJreiResult, err := fed.GetRecommenderationJeri(1566457662, 1566457662, 3600, 7, fedRcJeri, false)
	if err != nil {
		t.Fatal(err)
		return
	}
	logger.Infof("JERI Result: %s", utils.InterfaceToString(fedJreiResult))
}

func TestFedermeter_Calculate(t *testing.T) {
	fedProviders := &FedProviders{}
	// Instance 1
	fedRcNode1 := &FedProvider{}
	fedRcNode1.Region = "US West (Oregon)"
	fedRcNode1.Instances = &FedInstances{}
	fedRcNode1.Instances.Nodename = "172-23-1-8"
	fedRcNode1.Instances.Instancetype = "t2.xlarge"
	fedRcNode1.Instances.Nodetype = "master"
	fedRcNode1.Instances.Operatingsystem = "Linux"
	fedRcNode1.Instances.Preinstalledsw = "NA"
	fedRcNode1.Instances.Instancenum = "1"
	fedRcNode1.Instances.Period = "1"
	fedRcNode1.Instances.Unit = "month"
	fedRcNode1Storage := &FedStorage{}
	fedRcNode1Storage.Unit = "month"
	fedRcNode1Storage.Volumetype = "General Purpose"
	fedRcNode1Storage.Storagesize = "50"
	fedRcNode1Storage.Storagenum = "1"
	fedRcNode1Storage.Period = "1"
	fedRcNode1.Storage = append(fedRcNode1.Storage, fedRcNode1Storage)

	// Instance 2
	fedRcNode2 := &FedProvider{}
	fedRcNode2.Region = "US West (Oregon)"
	fedRcNode2.Instances = &FedInstances{}
	fedRcNode2.Instances.Nodename = "172-23-1-65"
	fedRcNode2.Instances.Instancetype = "t2.xlarge"
	fedRcNode2.Instances.Nodetype = "worker"
	fedRcNode2.Instances.Operatingsystem = "Linux"
	fedRcNode2.Instances.Preinstalledsw = "NA"
	fedRcNode2.Instances.Instancenum = "1"
	fedRcNode2.Instances.Period = "1"
	fedRcNode2.Instances.Unit = "month"
	fedRcNode2Storage := &FedStorage{}
	fedRcNode2Storage.Unit = "month"
	fedRcNode2Storage.Volumetype = "General Purpose"
	fedRcNode2Storage.Storagesize = "50"
	fedRcNode2Storage.Storagenum = "1"
	fedRcNode2Storage.Period = "1"
	fedRcNode2.Storage = append(fedRcNode2.Storage, fedRcNode2Storage)

	// Instance 3
	fedRcNode3 := &FedProvider{}
	fedRcNode3.Region = "US West (Oregon)"
	fedRcNode3.Instances = &FedInstances{}
	fedRcNode3.Instances.Nodename = "172-23-1-60"
	fedRcNode3.Instances.Instancetype = "t2.xlarge"
	fedRcNode3.Instances.Nodetype = "worker"
	fedRcNode3.Instances.Operatingsystem = "Linux"
	fedRcNode3.Instances.Preinstalledsw = "NA"
	fedRcNode3.Instances.Instancenum = "1"
	fedRcNode3.Instances.Period = "1"
	fedRcNode3.Instances.Unit = "month"
	fedRcNode3Storage := &FedStorage{}
	fedRcNode3Storage.Unit = "month"
	fedRcNode3Storage.Volumetype = "General Purpose"
	fedRcNode3Storage.Storagesize = "50"
	fedRcNode3Storage.Storagenum = "1"
	fedRcNode3Storage.Period = "1"
	fedRcNode3.Storage = append(fedRcNode3.Storage, fedRcNode3Storage)

	awsProvider := make(map[string][]*FedProvider, 0)
	awsProvider["aws"] = append(awsProvider["aws"], fedRcNode1)
	awsProvider["aws"] = append(awsProvider["aws"], fedRcNode2)
	awsProvider["aws"] = append(awsProvider["aws"], fedRcNode3)

	fedProviders.Calculator = append(fedProviders.Calculator, awsProvider)
	fmt.Println(utils.InterfaceToString(fedProviders))
	fedCalculateResult, err := fed.Calculate(fedProviders)
	if err != nil {
		t.Fatal(err)
		return
	}
	logger.Infof("Calculate Result: %s", utils.InterfaceToString(fedCalculateResult))

}