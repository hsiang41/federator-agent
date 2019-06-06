package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	CommonLib "github.com/containers-ai/api/common"
	"time"
	"math/rand"
	"fmt"
)

func GetTimeRange(startTime *timestamp.Timestamp, endTime *timestamp.Timestamp, durationTime int32, isInit bool, initGronularitySec int32) *CommonLib.TimeRange {
	if startTime == nil {
		startTime = ptypes.TimestampNow()
	}

	if durationTime == 0 {
		durationTime = 3600
	}

	if endTime == nil {
		eTm, _ := ptypes.Timestamp(startTime)
		eTm = eTm.Add(time.Duration(rand.Int31n(durationTime)) * time.Second)
		endTime, _ = ptypes.TimestampProto(eTm)
	}

	duTime, _ := time.ParseDuration(fmt.Sprintf("%ds", durationTime))
	du := ptypes.DurationProto(duTime)

	return &CommonLib.TimeRange{StartTime:startTime, EndTime:endTime, Step: du, AggregateFunction: 1}
}
