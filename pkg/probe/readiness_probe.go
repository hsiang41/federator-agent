package probe

import (
	"context"
	"fmt"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"google.golang.org/grpc"
)

type ReadinessProbeConfig struct {
	DataHubAddress string
	DataHubPort    int
}

func probeDataPipe(address string, port int) error {
	url := fmt.Sprintf("%s:%d", address, port)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := v1alpha1.NewDatahubServiceClient(conn)
	_, err = c.ListAlamedaNodes(context.Background(), &v1alpha1.ListAlamedaNodesRequest{})
	if err != nil {
		return err
	}

	return err
}
