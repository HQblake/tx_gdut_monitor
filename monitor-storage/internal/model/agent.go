package model

type AgentInfo struct {
	ID      int      `json:"id"`
	IP      string   `json:"ip"`
	Port    string   `json:"port"`
	Local   string   `json:"local"`
	IsLive  bool     `json:"isLive"`
	Metrics []string `json:"metrics"`
}
