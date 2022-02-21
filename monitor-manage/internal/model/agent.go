/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-13 14:39:53
 * @LastEditors: yzq
 */
package model

type AgentInfo struct {
	IP     string   `json:"ip"`
	Port   string   `json:"port"`
	Local  string   `json:"local"`
	IsLive bool     `json:"is_live"`
	Metric []string `json:"metric"`
}
type AgentSendInfo struct {
	IP         string  `json:"ip"`
	Port       string  `json:"port"`
	Local      string  `json:"local"`
	IsLive     bool    `json:"is_live"`
	SendConfig []int32 `json:"send"`
}
type Metrics struct {
	Metric []string `json:"metric"`
}
