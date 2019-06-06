package utils

import (
	"fmt"
	"testing"
	"time"
	"github.com/golang/protobuf/ptypes"
)

func TestGetTimeRange(t *testing.T) {
	sT := time.Now().UTC()
	sTimeStamp, _ := ptypes.TimestampProto(sT)
	tR := GetTimeRange(nil, sTimeStamp, 3600, false, 300)
	fmt.Println(tR.StartTime, tR.EndTime)
}