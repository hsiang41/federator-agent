package datahub

import (
	"testing"
	"fmt"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/google/uuid"
	datahubV1a1pha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"os"
)

var logger *log.Scope

func init() {
	logger = log.RegisterScope("datahubtest", "datahub testing", 0)
}

func TestDataHubClient_GetNodes(t *testing.T) {
	dp := NewDataHubClient()
	dp.Scope = logger
	dp.SetDataPipeAddress("52.12.208.57:30050")
	node, err := dp.GetNodes()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataHubClient_SendNotifyEvent(t *testing.T) {
	dp := NewDataHubClient()
	dp.Scope = logger
	dp.SetDataPipeAddress("127.0.0.1:50050")
	hostName, _ := os.Hostname()
	souce := &datahubV1a1pha1.EventSource{Host:hostName, Component: "federatorai-agent"}
	err := dp.SendNotifyEvent(
		uuid.New().String(), "", souce, datahubV1a1pha1.EventType_EVENT_TYPE_EMAIL_NOTIFICATION,
		datahubV1a1pha1.EventLevel_EVENT_LEVEL_ERROR, nil, "Test datahub client notify event",
		"Test datahub client notify event")
	if err != nil {
		t.Fatal(err)
	}
}