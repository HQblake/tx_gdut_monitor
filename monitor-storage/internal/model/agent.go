package model

type AgentInfo struct {
	Ip     string   `json:"ip"`
	Port   string   `json:"port"`
	Local  string   `json:"local"`
	IsLive bool     `json:"is_live"`
	Metric []string `json:"metric"`
}
