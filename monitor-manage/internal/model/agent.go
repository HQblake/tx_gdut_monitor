package model

type AgentInfo struct {
	Ip string `json:"ip"`
	Port string `json:"port"`
	Local string `json:"local"`
	IsLive string `json:"is_live"`
	Metric []string `json:"metric"`
}
type Metrics struct {
	Metric []string `json:"metric"`
}
