package model

type HistoryInfo struct {
	Id        int32   `json:"id"`
	Ip        string  `json:"ip"`
	Local     string  `json:"local"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Method    int32   `json:"method"`
	Level     int32   `json:"level"`
	Start     int64   `json:"start"`
	Duration  string  `json:"duration"`
}
