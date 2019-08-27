package datahub

import (
	"testing"
	"fmt"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

var logger *log.Scope

func init() {
	logger = log.RegisterScope("datahubtest", "datahub testing", 0)
}

func TestDataHubClient_GetNodes(t *testing.T) {
	dp := NewDataHubClient()
	dp.Scope = logger
	dp.SetDataPipeAddress("54.218.143.157:30946")
	node, err := dp.GetNodes()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.InterfaceToString(node))
}