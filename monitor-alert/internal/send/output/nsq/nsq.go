package nsq

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"github.com/nsqio/go-nsq"
)

type Nsq struct {
	level      output.Level
	formatType string
	producer  *nsq.Producer
	topic string
	address string
}

func NewNsq(level output.Level, config *Config) (*Nsq, error) {
	//创建生产者
	producer, err := createProducer(config.Address)
	if err != nil {
		return nil, err
	}
	return &Nsq{
		level: level,
		formatType: config.FormatType,
		topic: config.Topic,
		address: config.Address,
		producer: producer,
	}, nil
}


func (n *Nsq) Level() output.Level {
	return n.level
}

func (n *Nsq) Reset(level output.Level, config interface{}) error {
	conf, ok := config.(*Config)
	if !ok {
		return fmt.Errorf("config type is invalid")
	}
	err := conf.doCheck()
	if err != nil {
		return err
	}
	n.level = level
	n.formatType = conf.FormatType
	n.topic = conf.Topic
	// 地址有更新
	if conf.Address != n.address {
		if n.producer != nil {
			n.producer.Stop()
		}
		producer, err := createProducer(conf.Address)
		if err != nil {
			return err
		}
		n.producer = producer
	}
	return nil
}

func (n *Nsq) Output(info model.Info) error {
	msg, err := format.Format(n.formatType, info)
	if err != nil {
		return err
	}
	err = n.producer.Publish(n.topic, msg)
	if err != nil {
		return err
	}
	return nil
}

func (n *Nsq) Finish() error {
	n.producer.Stop()
	n.producer = nil
	return nil
}


func createProducer(address string) (*nsq.Producer, error) {
	nsqConf := nsq.NewConfig()
	producer, err := nsq.NewProducer(address, nsqConf)
	if err != nil {
		return nil, err
	}
	// 连通性测试
	err = producer.Ping()
	if err != nil {
		producer.Stop()
		producer = nil
		return nil, err
	}
	return producer, nil
}
