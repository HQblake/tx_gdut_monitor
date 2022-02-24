package nsq

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type Nsq struct {
	level      output.Level
	formatType string
	producer   *nsq.Producer
	topic      string
	address    string
	lock       *sync.RWMutex
}

func NewNsq(level output.Level, config *Config) (*Nsq, error) {
	//创建生产者
	producer, err := createProducer(config.Address)
	if err != nil {
		return nil, err
	}
	return &Nsq{
		lock:       &sync.RWMutex{},
		level:      level,
		formatType: config.FormatType,
		topic:      config.Topic,
		address:    config.Address,
		producer:   producer,
	}, nil
}

func (n *Nsq) Level() output.Level {
	n.lock.RLock()
	defer n.lock.RUnlock()
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
	n.lock.Lock()
	defer n.lock.Unlock()
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

func (n *Nsq) Output(infos []model.Info) error {
	msg, err := format.Format(n.formatType, infos)
	if err != nil {
		return err
	}
	n.lock.RLock()
	defer n.lock.RUnlock()
	if n.producer == nil {
		return nil
	}
	// 是否判定联通待定
	log.Println("nsq 告警开始")
	err = n.producer.Publish(n.topic, msg)
	if err != nil {
		return err
	}
	log.Println("nsq 告警结束")
	return nil
}

func (n *Nsq) Finish() error {
	n.lock.Lock()
	defer n.lock.Unlock()
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
