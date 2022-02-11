package configs

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"strings"
)

var cfg = new(config)

type config struct {
	AdminConfig *adminConfig `json:"admin" yaml:"admin"`
	StoreConfig *listenConfig `json:"store" yaml:"store"`
	AlertConfig *listenConfig `json:"alert" yaml:"alert"`
	DefaultRule *ruleConfig   `json:"rule" yaml:"rule"`
}
type adminConfig struct {
	Listen string `json:"listen" yaml:"listen"`
	ServerAddr string `json:"server" yaml:"server"`
}

type listenConfig struct {
	Listen string `json:"listen" yaml:"listen"`
}
type ruleConfig struct {
	Method int32 `json:"method" yaml:"method"`
	Period string `json:"period" yaml:"period"`
	Threshold map[string]float64 `json:"threshold" yaml:"threshold"`
}

func InitConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, cfg)
}

func GetDefaultConfig() *config {
	return cfg
}

func GetAdminListenAddr() string {
	return cfg.AdminConfig.Listen
}
func GetAdminServerAddr() string {
	return cfg.AdminConfig.ServerAddr
}

func GetStoreConnAddr() string {
	return cfg.StoreConfig.Listen
}

func GetAlertConnAddr() string {
	return cfg.AlertConfig.Listen
}

func GetDefaultRule() *defaultRule {
	return cfg.DefaultRule.parse()
}

func (r *ruleConfig) parse() *defaultRule {
	d := &defaultRule{
		Method: r.Method,
		Period: r.Period,
		Threshold: make(map[int32]float64),
	}
	for s, f := range r.Threshold {
		d.Threshold[parseLevel(s)] = f
	}
	return d
}


type defaultRule struct {
	Method int32 `json:"method" yaml:"method"`
	Period string `json:"period" yaml:"period"`
	Threshold map[int32]float64 `json:"threshold" yaml:"threshold"`
}

func parseLevel(lvl string) int32 {
	switch strings.ToLower(lvl) {
	case "panic":
		return 3
	case "error":
		return 2
	case "warn", "warning":
		return 1
	case "info":
		return 0
	}
	return 0
}