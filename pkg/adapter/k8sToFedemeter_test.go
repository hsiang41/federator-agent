package adapterK8sToFedmeter

import (
	"testing"
	"encoding/json"

	datahubV1a1pha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"fmt"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
)

func TestAdapterNodes_GenerateFedemeterNoes(t *testing.T) {
	var lsNodesDataHub datahubV1a1pha1.ListNodesResponse
	err := json.Unmarshal([]byte(nodes), &lsNodesDataHub)
	if err != nil {
		t.Fatal(err)
	}
	fed := NewAdapterNodes(lsNodesDataHub.Nodes)
	fedNodes, err := fed.GenerateFedemeterCalculates("month")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.InterfaceToString(fedNodes))
}

func TestAdapterNodes_GenerateFedemeterRecommendationNodes(t *testing.T) {
	var lsNodesDataHub datahubV1a1pha1.ListNodesResponse
	err := json.Unmarshal([]byte(nodes), &lsNodesDataHub)
	if err != nil {
		t.Fatal(err)
	}
	fed := NewAdapterNodes(lsNodesDataHub.Nodes)
	fedRecNodes, err := fed.GenerateFedemeterRecommendationNodes("jeri", "hour")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.InterfaceToString(fedRecNodes))
}


var nodes = `{"status":{},"nodes":[{"name":"ip-172-23-1-112.us-west-2.compute.internal","capacity":{"cpu_cores":2,"memory_bytes":4135084032},"provider":{"provider":"aws","instance_type":"t2.medium","region":"us-west-2","zone":"us-west-2a","os":"linux","role":"master","instance_id":"i-0769ec8570198bf4b","storage_size":46779129369}},{"name":"ip-172-23-1-67.us-west-2.compute.internal","capacity":{"cpu_cores":4,"memory_bytes":7834464256},"provider":{"provider":"aws","instance_type":"c4.xlarge","region":"us-west-2","zone":"us-west-2a","os":"linux","role":"worker","instance_id":"i-01a08072ceabff4ab","storage_size":46779129369}},{"name":"ip-172-23-1-63.us-west-2.compute.internal","capacity":{"cpu_cores":4,"memory_bytes":7834464256},"provider":{"provider":"aws","instance_type":"c4.xlarge","region":"us-west-2","zone":"us-west-2a","os":"linux","role":"worker","instance_id":"i-0364d9b65e9d1f973","storage_size":46779129369}}]}
`