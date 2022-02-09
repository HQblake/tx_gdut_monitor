package model


type JudgmentConfig struct {
	ID int32 `json:"id"`
	IP string `json:"ip"`
	Local string `json:"local"`
	Metric string `json:"metric"`
	Method int32 `json:"method"`
	Period string `json:"period"`
	Threshold string `json:"threshold"`
}