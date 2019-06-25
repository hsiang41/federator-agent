package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	CommonLib "github.com/containers-ai/api/common"
	"time"
	"fmt"
)

func GetTimeRange(startTime *timestamp.Timestamp, endTime *timestamp.Timestamp, durationTime int32, isInit bool, initGronularitySec int32) *CommonLib.TimeRange {
	if startTime == nil {
		sTm, _ := ptypes.Timestamp(endTime)
		du, _ := time.ParseDuration(fmt.Sprintf("-%ds", durationTime))
		sTm = sTm.Add(du)
		startTime, _ = ptypes.TimestampProto(sTm)
	}

	if durationTime == 0 {
		durationTime = 3600
	}

	if endTime == nil {
		eTm, _ := ptypes.Timestamp(startTime)
		du, _ := time.ParseDuration(fmt.Sprintf("%ds", durationTime))
		eTm = eTm.Add(du)
		endTime, _ = ptypes.TimestampProto(eTm)
	}

	duTime, _ := time.ParseDuration(fmt.Sprintf("%ds", initGronularitySec))
	du := ptypes.DurationProto(duTime)

	return &CommonLib.TimeRange{StartTime:startTime, EndTime:endTime, Step: du, AggregateFunction: 1}
}
