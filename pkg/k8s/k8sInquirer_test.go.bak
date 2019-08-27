package k8sInquirer

import (
	"fmt"
	"testing"

	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	Utils "github.com/containers-ai/federatorai-agent/pkg/utils"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"k8s.io/api/core/v1"
)

var logger *logUtil.Scope

func init () {
	logger = logUtil.RegisterScope("k8sInquirer", "k8s inquire", 0)
}

func TestK8sInquirer_GetNodes(t *testing.T) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf("unable to get k8s configuration, err: %v", err)
		return
	}

	k8sInq := Newk8sInquirer(cfg, logger)
	nodes, err := k8sInq.GetNodes()
	if err != nil {
		t.Errorf("Unable to list nodes, %v", err)
		return
	}
	fmt.Println("nodes:", Utils.InterfaceToString(nodes))
}

func TestK8sInquirer_GetPods(t *testing.T) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf("unable to get k8s configuration, err: %v", err)
		return
	}

	k8sInq := Newk8sInquirer(cfg, logger)
	pods, err := k8sInq.GetPods()
	if err != nil {
		t.Errorf("Unable to list pods, %v", err)
		return
	}
	fmt.Println("pods:", pods)
	for _, n := range pods.Items {
		logger.Infof("pod: %s", n.Name)
	}
}

func TestConvertK8sNode(t *testing.T) {
	var nodes v1.NodeList
	dat, err := ioutil.ReadFile("nodes.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal(dat, &nodes)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range nodes.Items{
		fmt.Println(n.Name)
		k8snode, err := ConvertK8sNode(&n)
		if err != nil {
			t.Fatal(err)
			continue
		}
		fmt.Printf("node: %s\n", Utils.InterfaceToString(k8snode))
	}
}