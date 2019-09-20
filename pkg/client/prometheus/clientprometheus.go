package ClientPrometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/prometheus/common/model"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	api "github.com/prometheus/client_golang/api/prometheus"
)

type MethodInt int

const (
	MethodQuery         MethodInt = 0
	MethodQueryRange	MethodInt = 1
	MethodSeries        MethodInt = 2
	MethodLabels        MethodInt = 3
	MethodLabelValues   MethodInt = 4
)

type PVector struct {
	Time        time.Time
	Metric      string
	Value       string
}

type ClientPrometheus struct {
	Addr    string
	Method  MethodInt
	Expr    string
	TimeRange *utils.TimeRange
}

func NewClientPrometheus(addr string, method MethodInt, expr string, timeRange *utils.TimeRange) *ClientPrometheus {
	return &ClientPrometheus{Addr: addr, Method: method, Expr: expr, TimeRange: timeRange}
}

func (c *ClientPrometheus) Execute() (string, error) {
	config := api.Config{Address: c.Addr}
	client, err := api.New(config)
	if err != nil {
		return "", err
	}
	switch c.Method {
	case MethodQuery:
		resp, err := query(&client, c.Expr, time.Now())
		if err != nil {
			return "", err
		}
		return resp, nil
	case MethodQueryRange:
		resp, err := queryRange(&client, c.Expr, *c.TimeRange)
		if err != nil {
			return "", err
		}
		return resp, nil
	}
	return "", status.Error(codes.Unimplemented, "Not supported")
}

func converJson(i model.Value) string {
	switch i.Type() {
	case model.ValNone:
		return "None"
	case model.ValScalar:
		v, _ := i.(*model.Scalar)
		result, _ := v.MarshalJSON()
		return string(result)
	case model.ValVector:
		var pVectors []PVector
		v, _ := i.(model.Vector)
		for _, m := range v {
			p := PVector{m.Timestamp.Time(), m.Metric.String(), m.Value.String()}
			pVectors = append(pVectors, p)
		}
		if len(pVectors) > 0 {
			result, err := json.Marshal(pVectors)
			if err != nil {
				return fmt.Sprintf("{\"error\": %s)", err.Error())
			}
			return string(result)
		}
		return ""
	case model.ValMatrix:
		var pVectors []PVector
		v, _ := i.(model.Matrix)
		for _, m := range v {
			for _, j := range m.Values {
				p := PVector{j.Timestamp.Time(), m.Metric.String(), j.Value.String()}
				pVectors = append(pVectors, p)
			}
		}
		if len(pVectors) > 0 {
			result, err := json.Marshal(pVectors)
			if err != nil {
				return fmt.Sprintf("{\"error\": %s)", err.Error())
			}
			return string(result)
		}
		return ""
	case model.ValString:
		result := i.String()
		return string(result)
	}
	return fmt.Sprintf("Unknown Type: %s", i.Type())
}

func query(c *api.Client, expr string, ts time.Time) (string, error) {
	queryClient := api.NewQueryAPI(*c)
	ctx := context.TODO()

	resp, err := queryClient.Query(ctx, expr, ts)
	if err != nil {
		return "", err
	}
	return converJson(resp), nil
}

func queryRange(c *api.Client, expr string, tr utils.TimeRange) (string, error) {
	queryClient := api.NewQueryAPI(*c)
	ctx := context.TODO()

	tmRange := api.Range{Start: tr.StartTime, End: tr.EndTime, Step: tr.Step}
	resp, err := queryClient.QueryRange(ctx, expr, tmRange)
	if err != nil {
		return "", err
	}

	return converJson(resp), nil
}
