package configs

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

var cfg = new(config)

type config struct {
	AdminConfig *adminConfig `json:"admin" yaml:"admin"`
}

type adminConfig struct {
	Listen string `json:"listen" yaml:"listen"`
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

