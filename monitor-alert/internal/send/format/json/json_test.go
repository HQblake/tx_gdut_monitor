package json

import (
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"reflect"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	cases := []struct{
		name string
		info model.Info
	}{
		{
			name: "full",
			info: model.Info{
				Agent:     "127.0.0.1 广州",
				Metric:    "CPU",
				Value:     10,
				Threshold: 3.14,
				Method:    "Sum",
				Level:     "panic",
				Duration:  "5min",
				Start:     time.Now().Format("[2006-01-01 15:04:05]"),
			},
		},
		{
			name: "section",
			info: model.Info{
				Agent:     "127.0.0.1 上海",
				Metric:    "memory",
				Value:     10,
				Threshold: 3.14,
				Method:    "Sum",
				Level:     "panic",
			},
		},
	}
	format := NewJson()
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			b, err := format.Format(cc.info)
			if err != nil {
				t.Errorf("format error:%v", err)
				return
			}
			got := model.Info{}
			err = json.Unmarshal(b, &got)
			if err != nil {
				t.Errorf("format %s, json unmarshal error:%v", string(b), err)
				return
			}
			t.Log(string(b))
			if !reflect.DeepEqual(got, cc.info) {
				t.Errorf("format error, want:%+v, got:%+v", cc.info, got)
			}
		})
	}
}
