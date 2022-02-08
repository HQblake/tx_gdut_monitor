package model

type CheckConfig struct {
	Id        int    `json:"id"`
	Ip        string `json:"ip"`
	Local     string `json:"local"`
	Metric    string `json:"metric"`
	Method    int    `json:"method"`
	Period    string `json:"period"`
	Threshold string `json:"threshold"`
}
