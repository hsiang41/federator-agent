package InfluxConvert

import (
	"testing"
	"github.com/containers-ai/api/datapipe/rawdata"
	v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"encoding/json"
	"fmt"
	"github.com/containers-ai/api/common"
)

func TestGetWriteRequest(t *testing.T) {
	var rawResponse rawdata.ReadRawdataResponse
	var fields []*InfluxField
	err := json.Unmarshal([]byte(upRawData), &rawResponse.Rawdata)
	if err != nil {
		t.Fatalf("Unable to parse up rawdata, %v", err)
	}
	iValue := &InfluxField{"value", common.DataType_DATATYPE_INT, ""}
	fields = append(fields, iValue)
	influxCV := NewInflux("alameda_prometheus", "cmd_api_servers_up", nil, fields, &rawResponse)
	rawWriteData, err := influxCV.GetWriteRequest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rawWriteData)
}

func TestGetWriteRequest_KubePodContainerStatusRestartsTotal(t *testing.T) {
	var rawResponse v1alpha1.ReadRawdataResponse
	var fields []*InfluxField
	tags := []string {"container", "instance", "namespace", "pod"}
	err := json.Unmarshal([]byte(kubePodContainerStatusRestartsTotal), &rawResponse)
	if err != nil {
		t.Fatalf("Unable to parse kube pods status rawdata, %v", err)
	}
	iValue := &InfluxField{"value", common.DataType_DATATYPE_INT, ""}
	fields = append(fields, iValue)
	influxCV := NewInflux("alameda_prometheus", "kube_pod_container_status_restarts_total", tags, fields, &rawResponse)
	rawWriteData, err := influxCV.GetWriteRequest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(rawWriteData.GetRawdata()))
	fmt.Println(rawWriteData.GetRawdata()[0].Columns)
	fmt.Println(rawWriteData.GetRawdata()[0].ColumnTypes)
	fmt.Println(rawWriteData.GetRawdata()[0].DataTypes)
}

var upRawData = `[
  {
    "query": {
      "table": "(sum(up{job=~\".*apiserver.*\"} == 1) / count(up{job=~\".*apiserver.*\"}))",
      "expression": "query_range",
      "condition": {
        "time_range": {
          "start_time": {
            "seconds": 1563854020,
            "nanos": 27083382
          },
          "end_time": {
            "seconds": 1563854035,
            "nanos": 27083382
          },
          "step": {
            "seconds": 60
          }
        },
        "order": 1,
        "limit": 1
      }
    },
    "columns": [
      "value"
    ],
    "groups": [
      {
        "rows": [
          {
            "time": {
              "seconds": 1563854020
            },
            "values": [
              "1"
            ]
          }
        ]
      }
    ],
    "rawdata": "{\"resultType\":\"matrix\",\"result\":[{\"metric\":{},\"values\":[[1563854020,\"1\"]]}]}"
  }
]`

var kubePodContainerStatusRestartsTotal = `{"status":{},"rawdata":[{"query":{"table":"kube_pod_container_status_restarts_total","expression":"query_range","condition":{"time_range":{"start_time":{"seconds":1563868091,"nanos":6790103},"end_time":{"seconds":1563868106,"nanos":6790103},"step":{"seconds":60}},"order":1,"limit":1}},"columns":["__name__","container","endpoint","instance","job","namespace","pod","service","value"],"groups":[{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","admission-controller","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","admission-controller-8784d7545-v6kb6","kube-state-metrics","3018"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alameda-datapipe","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","alameda-datapipe-7477d68964-wpfpc","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alameda-evictioner","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","alameda-evictioner-7f79dcb548-4jzp7","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alameda-executor","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","alameda-executor-697bc5d68b-j92fh","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alameda-operator","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","alameda-operator-5f98448d88-p4bf8","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-0","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-2","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-0","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","alertmanager-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-2","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","api","https-main","10.128.0.219:8443","kube-state-metrics","kube-system","master-api-oc-2-62","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","apiserver","https-main","10.128.0.219:8443","kube-state-metrics","kube-service-catalog","apiserver-7ml7c","kube-state-metrics","5"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","busybox","https-main","10.128.0.219:8443","kube-state-metrics","test","busybox-7c76c5f9c7-rgjkh","kube-state-metrics","1"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","cluster-monitoring-operator","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","cluster-monitoring-operator-6465f8fbc7-pqfvh","kube-state-metrics","8"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","config-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-0","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","config-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","config-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","alertmanager-main-2","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","console","https-main","10.128.0.219:8443","kube-state-metrics","openshift-console","console-67d46d959b-hln5z","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","controller-manager","https-main","10.128.0.219:8443","kube-state-metrics","kube-service-catalog","controller-manager-nfl6n","kube-state-metrics","12"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","controllers","https-main","10.128.0.219:8443","kube-state-metrics","kube-system","master-controllers-oc-2-62","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","etcd","https-main","10.128.0.219:8443","kube-state-metrics","kube-system","master-etcd-oc-2-62","kube-state-metrics","12"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","federatorai-agent","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","federatorai-agent-64c958c768-hkxbk","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","federatorai-portal","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","federatorai-portal-7887459d96-p56ns","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","grafana","https-main","10.128.0.219:8443","kube-state-metrics","federatorai","alameda-grafana-f78767845-jz25c","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","grafana","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","grafana-6b9f85786f-twgv4","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","grafana-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","grafana-6b9f85786f-twgv4","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","kube-rbac-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","node-exporter-4cfxd","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","kube-rbac-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","node-exporter-6mkvd","kube-state-metrics","2"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","kube-rbac-proxy-main","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","kube-state-metrics-7449d589bc-7ps7j","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","kube-rbac-proxy-self","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","kube-state-metrics-7449d589bc-7ps7j","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","kube-state-metrics","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","kube-state-metrics-7449d589bc-7ps7j","kube-state-metrics","7"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","node-exporter","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","node-exporter-4cfxd","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","node-exporter","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","node-exporter-6mkvd","kube-state-metrics","2"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","openvswitch","https-main","10.128.0.219:8443","kube-state-metrics","openshift-sdn","ovs-2f98x","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","openvswitch","https-main","10.128.0.219:8443","kube-state-metrics","openshift-sdn","ovs-4jmt6","kube-state-metrics","2"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-0","kube-state-metrics","4"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-1","kube-state-metrics","4"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus-config-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-0","kube-state-metrics","4"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus-config-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus-operator","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-operator-6644b8cd54-vk999","kube-state-metrics","0"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-0","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","prometheus-proxy","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","registry","https-main","10.128.0.219:8443","kube-state-metrics","default","docker-registry-1-ktxg6","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","registry-console","https-main","10.128.0.219:8443","kube-state-metrics","default","registry-console-1-lznnz","kube-state-metrics","4"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","router","https-main","10.128.0.219:8443","kube-state-metrics","default","router-1-9dpbv","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","rules-configmap-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-0","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","rules-configmap-reloader","https-main","10.128.0.219:8443","kube-state-metrics","openshift-monitoring","prometheus-k8s-1","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","sdn","https-main","10.128.0.219:8443","kube-state-metrics","openshift-sdn","sdn-l4mjn","kube-state-metrics","2"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","sdn","https-main","10.128.0.219:8443","kube-state-metrics","openshift-sdn","sdn-p5hpv","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","sync","https-main","10.128.0.219:8443","kube-state-metrics","openshift-node","sync-qvsdw","kube-state-metrics","3"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","sync","https-main","10.128.0.219:8443","kube-state-metrics","openshift-node","sync-tjh94","kube-state-metrics","2"]}]},{"rows":[{"time":{"seconds":1563868091},"values":["kube_pod_container_status_restarts_total","webconsole","https-main","10.128.0.219:8443","kube-state-metrics","openshift-web-console","webconsole-7df4f9f689-ntxts","kube-state-metrics","3"]}]}],"rawdata":"{\"resultType\":\"matrix\",\"result\":[{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"admission-controller\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"admission-controller-8784d7545-v6kb6\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3018\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alameda-datapipe\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"alameda-datapipe-7477d68964-wpfpc\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alameda-evictioner\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"alameda-evictioner-7f79dcb548-4jzp7\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alameda-executor\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"alameda-executor-697bc5d68b-j92fh\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alameda-operator\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"alameda-operator-5f98448d88-p4bf8\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-2\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"alertmanager-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-2\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"api\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"kube-system\",\"pod\":\"master-api-oc-2-62\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"apiserver\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"kube-service-catalog\",\"pod\":\"apiserver-7ml7c\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"5\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"busybox\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"test\",\"pod\":\"busybox-7c76c5f9c7-rgjkh\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"1\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"cluster-monitoring-operator\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"cluster-monitoring-operator-6465f8fbc7-pqfvh\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"8\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"config-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"config-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"config-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"alertmanager-main-2\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"console\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-console\",\"pod\":\"console-67d46d959b-hln5z\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"controller-manager\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"kube-service-catalog\",\"pod\":\"controller-manager-nfl6n\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"12\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"controllers\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"kube-system\",\"pod\":\"master-controllers-oc-2-62\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"etcd\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"kube-system\",\"pod\":\"master-etcd-oc-2-62\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"12\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"federatorai-agent\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"federatorai-agent-64c958c768-hkxbk\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"federatorai-portal\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"federatorai-portal-7887459d96-p56ns\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"grafana\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"federatorai\",\"pod\":\"alameda-grafana-f78767845-jz25c\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"grafana\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"grafana-6b9f85786f-twgv4\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"grafana-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"grafana-6b9f85786f-twgv4\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"kube-rbac-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"node-exporter-4cfxd\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"kube-rbac-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"node-exporter-6mkvd\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"2\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"kube-rbac-proxy-main\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"kube-state-metrics-7449d589bc-7ps7j\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"kube-rbac-proxy-self\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"kube-state-metrics-7449d589bc-7ps7j\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"kube-state-metrics\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"kube-state-metrics-7449d589bc-7ps7j\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"7\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"node-exporter\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"node-exporter-4cfxd\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"node-exporter\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"node-exporter-6mkvd\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"2\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"openvswitch\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-sdn\",\"pod\":\"ovs-2f98x\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"openvswitch\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-sdn\",\"pod\":\"ovs-4jmt6\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"2\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"4\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"4\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus-config-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"4\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus-config-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus-operator\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-operator-6644b8cd54-vk999\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"0\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"prometheus-proxy\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"registry\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"default\",\"pod\":\"docker-registry-1-ktxg6\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"registry-console\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"default\",\"pod\":\"registry-console-1-lznnz\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"4\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"router\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"default\",\"pod\":\"router-1-9dpbv\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"rules-configmap-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-0\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"rules-configmap-reloader\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-monitoring\",\"pod\":\"prometheus-k8s-1\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"sdn\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-sdn\",\"pod\":\"sdn-l4mjn\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"2\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"sdn\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-sdn\",\"pod\":\"sdn-p5hpv\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"sync\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-node\",\"pod\":\"sync-qvsdw\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"sync\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-node\",\"pod\":\"sync-tjh94\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"2\"]]},{\"metric\":{\"__name__\":\"kube_pod_container_status_restarts_total\",\"container\":\"webconsole\",\"endpoint\":\"https-main\",\"instance\":\"10.128.0.219:8443\",\"job\":\"kube-state-metrics\",\"namespace\":\"openshift-web-console\",\"pod\":\"webconsole-7df4f9f689-ntxts\",\"service\":\"kube-state-metrics\"},\"values\":[[1563868091,\"3\"]]}]}"}]}`