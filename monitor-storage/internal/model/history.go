package model

import "sync"

type HistoryInfo struct {
	ID        int32   `json:"id"`
	IP        string  `json:"ip"`
	Local     string  `json:"local"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Method    int32   `json:"method"`
	Level     int32   `json:"level"`
	Start     int64   `json:"start"`
	Duration  string  `json:"duration"`
}

var HistoryInfoPool = &sync.Pool{
	New: func() interface{} {
		return new(HistoryInfo)
	},
}
