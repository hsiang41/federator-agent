package ClientPrometheus

import (
	"testing"
	"fmt"
	"github.com/containers-ai/federatorai-agent/pkg/utils"
	"time"
)

func TestClientPrometheus_Execute_Query(t *testing.T) {
	pQuery := NewClientPrometheus("http://127.0.0.1:19090", MethodQuery, "nvidia_gpu_duty_cycle", nil)
	result, err := pQuery.Execute()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestClientPrometheus_Execute_Query_TimeRange(t *testing.T) {
	tr := &utils.TimeRange{}
	tr.EndTime = time.Now()
	td, _ := time.ParseDuration("-30s")
	tr.StartTime = tr.EndTime.Add(td)
	tr.Step, _ = time.ParseDuration("30s")
	pQuery := NewClientPrometheus("http://127.0.0.1:19090", MethodQueryRange, "nvidia_gpu_duty_cycle{minor_number=\"0\", uuid=\"GPU-4be3789d-3fea-f010-7f10-0f024f2d1afa\"}", tr)
	result, err := pQuery.Execute()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}