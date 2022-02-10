package model

type AgentInfo struct {
	IP string `json:"ip"`
	Port string `json:"port"`
	Local string `json:"local"`
	IsLive bool `json:"is_live"`
	Metric []string `json:"metric"`
}
type Metrics struct {
	Metric []string `json:"metric"`
}
