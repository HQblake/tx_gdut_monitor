package model


type SendConfig struct {
	Id int `json:"id"`
	Ip string `json:"ip"`
	Local string `json:"local"`
	SendType int `json:"send_type"`
	Level int `json:"level"`
	Config string `json:"config"`
}