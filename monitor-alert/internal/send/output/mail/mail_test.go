package mail

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/html"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"testing"
	"time"
)

var infos = []model.Info{
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
func TestSendMail(t *testing.T) {
	html.Register()
	cases := []struct{
		name string
		info []model.Info
		level output.Level
		conf *EMailConf
	}{
		{
			name: "one mail",
			info: []model.Info{infos[0]},
			conf : &EMailConf{
				Target: []string{"526756656@qq.com"},
				FormatType: "html",
			},
		},
		{
			name: "many mail",
			info: infos[:5],
			level: output.WarnLevel,
			conf : &EMailConf{
				Target: []string{"526756656@qq.com"},
				FormatType: "html",
			},
		},
		{
			name: "more mail",
			info: infos,
			level: output.WarnLevel,
			conf : &EMailConf{
				Target: []string{"526756656@qq.com","zekewcs@163.com"},
				FormatType: "html",
			},
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			outputs, err := NewMail(cc.level, cc.conf)
			if err != nil {
				t.Fatal(err)
			}

			err = outputs.Output(cc.info)
			if err != nil {
				t.Error(err)
			}
			err = outputs.Finish()
			if err != nil {
				t.Fatal(err)
			}

		})
	}
	log.Println("send successfully ... ")
}

func TestResetMail(t *testing.T) {
	html.Register()
	cases := []struct{
		name string
		info []model.Info
		level output.Level
		conf *Config
	}{
		{
			name: "one mail",
			info: []model.Info{infos[0]},
			conf : &Config{
				Target: "526756656@qq.com",
				FormatType: "html",
			},
		},
		{
			name: "many mail",
			info: infos[:5],
			level: output.WarnLevel,
			conf : &Config{
				Target: "526756656@qq.com",
				FormatType: "html",
			},
		},
		{
			name: "more mail",
			info: infos,
			level: output.WarnLevel,
			conf : &Config{
				Target: "526756656@qq.com,zekewcs@163.com",
				FormatType: "html",
			},
		},
	}
	outputs, err := NewMail(output.WarnLevel, &EMailConf{})
	if err != nil {
		t.Fatal(err)
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			err = outputs.Reset(cc.level, cc.conf)
			if err != nil {
				t.Fatal(err)
			}
			err = outputs.Output(cc.info)
			if err != nil {
				t.Error(err)
			}
		})
	}
	err = outputs.Finish()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("send successfully ... ")
}
