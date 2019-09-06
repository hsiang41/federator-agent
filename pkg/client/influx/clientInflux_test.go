package ClientInflux

import (
	"testing"
	"fmt"
)

func TestClientInflux_Execute(t *testing.T) {
	ifxdb := NewClientInflux("http://127.0.0.1:8086", "gpu_counting", MethodQuery, "select * from k8s where time > now() - 3600s order by time limit 10")
	result, err := ifxdb.Execute()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Sprintf(result)
}