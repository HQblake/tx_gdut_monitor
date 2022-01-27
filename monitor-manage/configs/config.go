package configs

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"manage/internal/database"
)

var cfg = new(config)

type config struct {
	AdminConfig *adminConfig `json:"admin" yaml:"admin"`
	DbConfig    *dbConfig    `json:"db" yaml:"db"`
}

type dbConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
	Type     string `json:"type" yaml:"type"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
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

func GetDataBaseConfig() database.Config {
	return cfg.DbConfig
}

func (d *dbConfig) GetDriver() string {
	return d.Type
}

func (d *dbConfig) GetSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", d.User, d.Password, d.Host, d.Port, d.Name)
}
