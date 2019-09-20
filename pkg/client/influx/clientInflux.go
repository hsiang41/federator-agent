package ClientInflux

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
)

type MethodInt int

const (
	MethodQuery MethodInt = 0
	MethodWrite MethodInt = 1
)

type ClientInflux struct {
	Addr    string
	Database string
	Method  MethodInt
	Expr    string
	TimeRange *utils.TimeRange
}

func NewClientInflux(addr string, database string, method MethodInt, expr string) *ClientInflux {
	return &ClientInflux{Addr: addr, Database: database, Method: method, Expr: expr}
}

func (n *ClientInflux) Execute() (string, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: n.Addr})
	if err != nil {
		return "", err
	}
	defer c.Close()
	switch n.Method {
	case MethodQuery:
		return query(&c, n.Database, n.Expr)
	case MethodWrite:
		return "", status.Error(codes.Unimplemented, "Not implemented")
	}
	return "", status.Error(codes.Unimplemented, "Not supported")
}

func query(c *client.Client, database string, expr string) (string, error) {
	q := client.NewQuery(expr, database, "")
	response, err := (*c).Query(q)
	if err != nil {
		return "", err
	}
	if response.Error() != nil {
		return "", response.Error()
	}
	return utils.InterfaceToString(response.Results), nil
}