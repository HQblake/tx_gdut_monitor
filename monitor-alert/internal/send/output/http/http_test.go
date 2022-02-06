package http

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSendHttp(t *testing.T) {
	json.Register()
	target := httptest.NewServer(http.HandlerFunc(f)).URL
	infos := []model.Info{
		{
			Agent: "127.0.0.1 广州",
			Metric: "CPU",
			Level:"panic",
			Duration:"5min",
			Start:time.Now().Format("[2006-01-01 15:04:05]"),
		},
		{
			Agent: "127.0.0.4 上海",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.1 北京",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.3 深圳",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.1 深圳",
			Metric: "cpu",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.2 深圳",
			Metric: "网络",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
	}
	cases := []struct{
		name string
		info []model.Info
		level output.Level
		conf *Config
	}{
		{
			name: "get test",
			info: []model.Info{infos[0]},
			conf : &Config{
				Method: http.MethodGet,
				Url: target,
				FormatType: "json",
				Headers: map[string]string{},
			},
		},
		{
			name: "post test",
			info: infos,
			conf : &Config{
				Method: http.MethodPost,
				Url: target,
				FormatType: "json",
				Headers: map[string]string{},
			},
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			outputs, err := NewHttp(cc.level, cc.conf)
			if err != nil {
				t.Fatal(err)
			}
			for _, info := range cc.info {
				err = outputs.Output(info)
				if err != nil {
					t.Error(err)
				}
			}
			err = outputs.Finish()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestResetHttp(t *testing.T) {
	json.Register()
	level := output.InfoLevel
	conf := &Config{}
	target := httptest.NewServer(http.HandlerFunc(f)).URL
	infos := []model.Info{
		{
			Agent: "127.0.0.1 广州",
			Metric: "CPU",
			Level:"panic",
			Duration:"5min",
			Start:time.Now().Format("[2006-01-01 15:04:05]"),
		},
		{
			Agent: "127.0.0.4 上海",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.1 北京",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.3 深圳",
			Metric: "内存",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.1 深圳",
			Metric: "cpu",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
		{
			Agent: "127.0.0.2 深圳",
			Metric: "网络",
			Value:10,
			Threshold:3.14,
			Method:"Sum",
		},
	}
	cases := []struct{
		name string
		info []model.Info
		level output.Level
		conf *Config
	}{
		{
			name: "get test",
			info: []model.Info{infos[0]},
			level: output.WarnLevel,
			conf : &Config{
				Method: http.MethodGet,
				Url: target,
				FormatType: "json",
				Headers: map[string]string{},
			},
		},
		{
			name: "post test",
			info: infos,
			level: output.WarnLevel,
			conf : &Config{
				Method: http.MethodPost,
				Url: target,
				FormatType: "json",
				Headers: map[string]string{},
			},
		},
	}
	outputs, err := NewHttp(level, conf)
	if err != nil {
		t.Fatal(err)
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			err = outputs.Reset(cc.level, cc.conf)
			if err != nil {
				t.Fatal(err)
			}
			for _, info := range cc.info {
				err = outputs.Output(info)
				if err != nil {
					t.Error(err)
				}
			}
		})
	}
	err = outputs.Finish()
	if err != nil {
		t.Fatal(err)
	}
}

func f(writer http.ResponseWriter, request *http.Request)  {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("http receive :", string(body))
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("receive successfully"))
}