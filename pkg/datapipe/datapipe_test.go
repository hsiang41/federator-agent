package datapipe

import (
	"testing"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/alameda/operator/datahub"
	"github.com/golang/mock/gomock"
)

func TestGetNodes(t *testing.T) {
	dp := NewDataPipeClient()
	dp.Scope = logUtil.RegisterScope("federatorai-agent", "federatorai-agent datapipe self test", 0)
	dp.DataPipe = DataPipeConfig{&datahub.Config{}, 30, 300, 3600 }
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
}