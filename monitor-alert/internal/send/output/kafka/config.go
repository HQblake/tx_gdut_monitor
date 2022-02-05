package kafka

import "strings"

type Config struct {
	FormatType string `json:"format_type"`
}

func (c *Config) doCheck() error {
	c.FormatType = strings.ToLower(c.FormatType)
	if c.FormatType == "" {
		c.FormatType = "line"
	}
	return nil
}
