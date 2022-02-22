package nsq

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format/line"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
	"time"
)

const (
	// 仅供临时测试
	address = "81.71.165.211:4150"
	topic   = "test"
	channel = "test_channel"
)

func TestPublish(t *testing.T) {
	line.Register()
	infos := []model.Info{
		{
			Agent:    "127.0.0.1 广州",
			Metric:   "CPU",
			Level:    "panic",
			Duration: "5min",
			Start:    time.Now().Format("[2006-01-01 15:04:05]"),
		},
		{
			Agent:     "127.0.0.4 上海",
			Metric:    "内存",
			Value:     10,
			Threshold: 3.14,
			Method:    "Sum",
		},
		{
			Agent:     "127.0.0.1 北京",
			Metric:    "内存",
			Value:     10,
			Threshold: 3.14,
			Method:    "Sum",
		},
		{
			Agent:     "127.0.0.3 深圳",
			Metric:    "内存",
			Value:     10,
			Threshold: 3.14,
			Method:    "Sum",
		},
		{
			Agent:     "127.0.0.1 深圳",
			Metric:    "cpu",
			Value:     10,
			Threshold: 3.14,
			Method:    "Sum",
		},
		{
			Agent:     "127.0.0.2 深圳",
			Metric:    "网络",
			Value:     10,
			Threshold: 3.14,
			Method:    "Sum",
		},
	}
	cases := []struct {
		name  string
		info  []model.Info
		level output.Level
		conf  *Config
	}{
		{
			name:  "single test",
			info:  infos,
			level: 0,
			conf: &Config{
				Topic:      topic,
				Address:    address,
				FormatType: "line",
			},
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			outputs, err := NewNsq(cc.level, cc.conf)
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

func TestMain(t *testing.M) {

	go createConsumer(address, topic, channel)
	t.Run()
	<-time.After(5 * time.Second)
}

// 消费者
type Consumer struct{}

func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}
func createConsumer(address string, topic string, channel string) {
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		log.Fatal(err.Error())
	}
	c.AddHandler(&Consumer{})
	if err = c.ConnectToNSQD(address); err != nil {
		log.Fatal(err.Error())
	}

}
