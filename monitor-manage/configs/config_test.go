package configs

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	target := &config{
		AdminConfig: &listenConfig{
			Listen: "8080",
		},
		DefaultRule: &ruleConfig{
			Method: 1,
			Period: "5m",
			Threshold: map[string]float64{
				"warn": 0.03,
				"error": 0.06,
				"panic": 0.08,
			},
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
	t.Log(GetDefaultRule())

}
