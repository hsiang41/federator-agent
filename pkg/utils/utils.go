package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	CommonLib "github.com/containers-ai/api/common"
	"time"
	"fmt"
	"encoding/json"
)

func GetTimeRange(startTime *timestamp.Timestamp, endTime *timestamp.Timestamp, durationTime int32, isInit bool, step int32) *CommonLib.TimeRange {
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

	duTime, _ := time.ParseDuration(fmt.Sprintf("%ds", step))
	du := ptypes.DurationProto(duTime)

	return &CommonLib.TimeRange{StartTime:startTime, EndTime:endTime, Step: du, AggregateFunction: 0}
}

func InterfaceToString(i interface{}, params ...string) string {
	var indent string
	var err error
	var v []byte
	if params != nil && len(params) > 0 {
		indent = params[0]
	}
	if len(indent) > 0 {
		v, err = json.MarshalIndent(i, "", indent)
	} else {
		v, err = json.Marshal(i)
	}
	if err != nil {
		return ""
	}
	return string(v)
}
