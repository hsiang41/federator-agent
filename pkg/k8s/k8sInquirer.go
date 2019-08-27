package k8sInquirer

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type k8sInquirer struct {
	client client.Client
	config *rest.Config
	logger *logUtil.Scope
	mgr    *manager.Manager
}

func Newk8sInquirer (config *rest.Config, logger *logUtil.Scope ) *k8sInquirer {
	cli, err := client.New(config, client.Options{})
	if err != nil {
		logger.Errorf("Unable to create client, %v", err)
		return nil
	}
	mgr, err := manager.New(config, manager.Options{})
	if err != nil {
		logger.Errorf("Unable to create manager, %v", err)
		return nil
	}
	return &k8sInquirer{config: config, client: cli, logger: logger, mgr: &mgr}
}

func (k *k8sInquirer) getNodes() (*corev1.NodeList, error) {
	nodes := &corev1.NodeList{}
	err := k.client.List(context.TODO(), &client.ListOptions{}, nodes)
	if err != nil {
		k.logger.Errorf("Unable to list nodes, %v", err)
		return nil, err
	}
	return nodes, nil
}

func parseKeyValue(strParse string, key string, value string) string {
	pattern, err := regexp.Compile(strings.ToLower(fmt.Sprintf("/%s$", key)))
	if err != nil {
		return ""
	}
	if len(pattern.FindString(strings.Replace(strParse, "-", "", -1))) > 0 {
		return value
	}
	return ""
}

func parseProviderID(providerID string) (string, string, string) {
	var provider string
	var region string
	var instanceID string
	rex, err := regexp.Compile("([^\\:/]+)")
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	res := rex.FindAllString(providerID, -1)
	if res == nil || len(res) == 0 {
		return "", "", ""
	}
	for i := 0; i < len(res) && i < 3; i++ {
		switch i {
		case 0:
			provider = res[i]
		case 1:
			region = res[i]
		case 2:
			instanceID = res[i]
		}
	}
	return provider, region, instanceID
}

func ConvertK8sNode(node *corev1.Node) (*k8sNode, error) {
	nodesLabel := &k8sNode{Name: node.Name, Namespaces: node.Namespace, Kind: node.Kind}
	rf := reflect.TypeOf(*nodesLabel)
	rv := reflect.ValueOf(nodesLabel).Elem()
	for i := 0; i < rf.NumField(); i++ {
		key := rf.Field(i).Name
		// parse node label information
		for labelKey, labelV := range node.Labels {
			if strings.Contains(labelKey, "stackpoint.") && strings.Contains(labelKey, "stackpoint.io/role") == false {
				continue
			}
			value := parseKeyValue(labelKey, key, labelV)
			if len(value) > 0 {
				rValue := rv.FieldByName(strings.Title(key))
				rValue.SetString(string(labelV))
				break
			}
		}
		switch key {
		case "StorageSize":
			nodesLabel.StorageSize = node.Status.Allocatable.StorageEphemeral().Value()
		case "Cpu":
			nodesLabel.Cpu = node.Status.Allocatable.Cpu().Value()
		}
	}

	if len(node.Spec.ProviderID) > 0 {
		provider, _, instanceID := parseProviderID(node.Spec.ProviderID)
		nodesLabel.Provider = provider
		nodesLabel.InstanceID = instanceID
	}
	return nodesLabel, nil
}

func (k *k8sInquirer) GetNodes() ([]*k8sNode, error) {
	lsK8sNode := make([]*k8sNode, 0)
	nodes, err := k.getNodes()
	if err != nil {
		return nil, err
	}
	for _, n := range nodes.Items {
		k8sNode, err := ConvertK8sNode(&n)
		if err != nil {
			continue
		}
		lsK8sNode = append(lsK8sNode, k8sNode)
	}
	return lsK8sNode, nil
}

func (k *k8sInquirer) GetPods() (*corev1.PodList, error) {
	pods := &corev1.PodList{}
	err := k.client.List(context.TODO(), &client.ListOptions{}, pods)
	if err != nil {
		k.logger.Errorf("Unable to list pods, %v", err)
		return nil, err
	}

	return pods, nil
}

// instance-typ --> size
var nodeLabels = `{      
      "beta.kubernetes.io/arch": "amd64",
      "beta.kubernetes.io/instance-type": "c4.xlarge",
      "beta.kubernetes.io/os": "linux",
      "failure-domain.beta.kubernetes.io/region": "us-west-2",
      "failure-domain.beta.kubernetes.io/zone": "us-west-2a",
      "kubernetes.io/arch": "amd64",
      "kubernetes.io/hostname": "ip-172-23-1-67.us-west-2.compute.internal",
      "kubernetes.io/os": "linux",
      "stackpoint.io/cluster_id": "23391",
      "stackpoint.io/instance_id": "netatt9dgn-worker-2",
      "stackpoint.io/node_group": "autoscaling-netatt9dgn-pool-1",
      "stackpoint.io/node_id": "91192",
      "stackpoint.io/node_pool": "Default-Worker-Pool",
      "stackpoint.io/private_ip": "172.23.1.67",
      "stackpoint.io/role": "worker",
      "stackpoint.io/size": "c4.xlarge"
}`

// providerID: aws:///us-west-2a/i-0769ec8570198bf4b --> <provider_raw>//<region>//<instance_id>
