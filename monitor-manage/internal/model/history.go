package model

type HistoryInfo struct {
	Id int `json:"id"`
	Ip string `json:"ip"`
	Local string `json:"local"`
	Metric string `json:"metric"`
	Value float64 `json:"value"`
	Threshold string `json:"threshold"`
	Method int `json:"method"`
	Level int `json:"level"`
	Start string `json:"start"`
	Duration string `json:"duration"`
}
