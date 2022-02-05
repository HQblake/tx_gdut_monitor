package mail

import (
	"fmt"
	"strings"
)

type Config struct {
	Target []string `json:"target"`
	FormatType string `json:"format_type"`
}

func (c *Config) doCheck() error {
	c.FormatType = strings.ToLower(c.FormatType)
	for _, s := range c.Target {
		if strings.Index(s, "@") == -1 {
			return fmt.Errorf("email %s is invalid", s)
		}
	}
	return nil
}

