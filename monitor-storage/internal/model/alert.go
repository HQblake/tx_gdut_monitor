package model

import "sync"

type AlertConfig struct {
	ID       int32  `json:"id"`
	IP       string `json:"ip"`
	Local    string `json:"local"`
	SendType int32  `json:"sendType"`
	Config   string `json:"config"`
	Level    int32  `json:"level"`
}

var AlertConfigPool = sync.Pool{New: func() interface{} {
	return new(AlertConfig)
}}
