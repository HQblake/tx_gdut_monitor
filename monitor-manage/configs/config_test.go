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
