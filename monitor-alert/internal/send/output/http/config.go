package http

import (
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	FormatType string            `json:"format_type"`
	Method     string            `json:"method"`
	Url        string            `json:"url"`
	Headers    map[string]string `json:"headers"`
}

func (c *Config) Check() error {
	return c.doCheck()
}

func (c *Config) doCheck() error {
	c.FormatType = strings.ToLower(c.FormatType)
	if c.FormatType == "" {
		c.FormatType = "json"
	}
	c.Method = strings.ToUpper(c.Method)
	switch c.Method {
	case "POST", "HEAD", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE":
	default:
		return fmt.Errorf("method is invalid")
	}
	if c.Url == "" {
		return fmt.Errorf("url can not be null")
	}
	if !strings.HasPrefix(c.Url, "http://") && !strings.HasPrefix(c.Url, "https://") {
		return fmt.Errorf("url is invalid")
	}
	return nil
}

func toHeader(h map[string]string) http.Header {
	header := make(http.Header)
	for k, v := range h {
		header.Set(k, v)
	}
	return header
}
