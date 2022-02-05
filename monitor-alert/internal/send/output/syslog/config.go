package syslog

type Config struct {
	FormatType string `json:"format_type"`
	Network string `json:"network"`
	Address string `json:"address"`
}

