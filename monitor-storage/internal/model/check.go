package model

import "sync"

type CheckConfig struct {
	ID        int32  `json:"id"`
	IP        string `json:"ip"`
	Local     string `json:"local"`
	Metric    string `json:"metric"`
	Method    int32  `json:"method"`
	Period    string `json:"period"`
	Threshold string `json:"threshold"`
}

var CheckConfigPool = sync.Pool{New: func() interface{} {
	return new(CheckConfig)
}}
