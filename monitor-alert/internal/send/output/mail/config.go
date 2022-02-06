package mail

import (
	"fmt"
	"strings"
)

type Config struct {
	Target string `json:"target"`
	FormatType string `json:"format_type"`
}

type EMailConf struct {
	Target []string `json:"target"`
	FormatType string `json:"format_type"`
}

func (c *Config) doCheck() (*EMailConf, error) {
	m := &EMailConf{}
	m.FormatType = strings.ToLower(c.FormatType)
	if m.FormatType == "" {
		m.FormatType = "html"
	}
	// 方便前端处理，以逗号分割
	addr := strings.Split(c.Target, ",")
	if len(addr) == 0 {
		return nil, fmt.Errorf("address can not be null")
	}
	for _, s := range addr {
		if strings.Index(s, "@") == -1 {
			return nil, fmt.Errorf("email %s is invalid", s)
		}
	}
	m.Target = addr
	return m, nil
}

