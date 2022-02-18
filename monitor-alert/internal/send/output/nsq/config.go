package nsq

import (
	"fmt"
	"strings"
)

type Config struct {
	FormatType string `json:"format_type"`
	Topic      string `json:"topic"`
	Address    string `json:"address"`
}

func (c *Config) Check() error {
	return c.doCheck()
}

func (c *Config) doCheck() error {
	c.FormatType = strings.ToLower(c.FormatType)
	if c.FormatType == "" {
		c.FormatType = "line"
	}
	if c.Topic  == ""{
		return fmt.Errorf("topic can not be null")
	}
	if c.Address  == ""{
		return fmt.Errorf("address can not be null")
	}
	return nil
}
