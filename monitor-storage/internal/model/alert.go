package model

import "sync"

// AlertInfo 记录某台Agent的所有异常指标信息
type AlertInfo struct {
	AgentID string
	Metrics map[string]MetricInfo
}

// MetricInfo 记录异常指标信息
type MetricInfo struct {
	ID        int32
	Metric    string
	Value     float64
	Threshold float64
	Method    int8
	Level     int8
	Duration  string
	Start     int64
}

var AlertInfoPool = &sync.Pool{
	New: func() interface{} {
		return new(AlertInfo)
	},
}

type HistoryInfo struct {
	Id        int     `json:"id"`
	Ip        string  `json:"ip"`
	Local     string  `json:"local"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold string  `json:"threshold"`
	Method    int     `json:"method"`
	Level     int     `json:"level"`
	Start     string  `json:"start"`
	Duration  string  `json:"duration"`
}
