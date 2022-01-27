package configs

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	target := &config{
		AdminConfig: &adminConfig{
			Listen: "8080",
		},
		DbConfig: &dbConfig{
			Host: "localhost",
			Port: "3306",
			Type: "mysql",
			Name: "tx_gdut_monitor",
			User: "root",
			Password: "123456",
		},
	}
	err := InitConfig("./config.yml")
	if err != nil {
		t.Errorf("init config error %v", err)
	}
	c := GetDefaultConfig()
	if !reflect.DeepEqual(c, target) {
		t.Error("read config error")
	}

}
