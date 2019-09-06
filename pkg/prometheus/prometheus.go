package prometheus

import (
	"fmt"
	"github.com/containers-ai/federatorai-agent/pkg/datapipe"
	"github.com/containers-ai/api/common"
	"github.com/containers-ai/api/datapipe/rawdata"
)

type PromQL struct {
	Expr string
	EndPoint string
	Fields []string
	DataSource *datapipe.DataPipeClient
}

func NewPromQL(expr string, endPoint string, fields []string, client *datapipe.DataPipeClient) *PromQL {
	if len(endPoint) == 0 {
		endPoint = "query_range"
	}
	return &PromQL{
		Expr: expr,
		EndPoint: endPoint,
		Fields: fields,
		DataSource: client}
}


func (p *PromQL) GetRawData(timeRange *common.TimeRange, limit uint64) (*rawdata.ReadRawdataResponse, error) {
	rawsData, err := p.DataSource.GetRawData(p.Expr, p.EndPoint, timeRange, limit)
	if err != nil {
		return nil, err
	}

	if rawsData.Status != nil && rawsData.Status.Code != 0 {
		return nil, fmt.Errorf("Failed to get prometheus raw data, status: %d, %s", rawsData.Status.GetCode(), rawsData.Status.GetMessage())
	}

	return rawsData, nil
}
