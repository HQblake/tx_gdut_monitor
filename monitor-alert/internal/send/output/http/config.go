package http

import "strings"

type Config struct {
	FormatType string `json:"format_type"`
	Method string `json:"method"`
	Url string `json:"url"`
	Headers map[string]string `json:"headers"`
	
}
func (c *Config) doCheck() error {
	c.FormatType = strings.ToLower(c.FormatType)
	if c.FormatType == "" {
		c.FormatType = "json"
	}
	return nil
}
