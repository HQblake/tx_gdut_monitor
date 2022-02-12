package model

type SendConfig struct {
	ID       int32  `json:"id"`
	IP       string `json:"ip"`
	Local    string `json:"local"`
	SendType int32  `json:"send_type"`
	Level    int32  `json:"level"`
	Config   string `json:"config"`
}
