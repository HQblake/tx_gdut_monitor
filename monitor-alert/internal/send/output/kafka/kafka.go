package kafka

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

type Kafka struct {
	level      output.Level
	formatType string
	conf       *ProducerConfig
	enable     bool
	input      chan<- *sarama.ProducerMessage
	producer   sarama.AsyncProducer
	lock       *sync.RWMutex
	wg         *sync.WaitGroup
	cancel     context.CancelFunc
}

func NewKafka(level output.Level, conf *ProducerConfig) (*Kafka, error) {
	producer, err := sarama.NewAsyncProducer(conf.Address, conf.Conf)
	if err != nil {
		return nil, err
	}
	k := &Kafka{
		lock:       &sync.RWMutex{},
		level:      level,
		conf:       conf,
		formatType: conf.FormatType,
		enable:     false,
		input:      producer.Input(),
		producer:   producer,
	}
	go k.sendMsg()
	return k, nil
}

func (k *Kafka) Level() output.Level {
	k.lock.RLock()
	defer k.lock.RUnlock()
	return k.level
}

func (k *Kafka) Reset(level output.Level, config interface{}) error {
	conf, ok := config.(*Config)
	if !ok {
		return fmt.Errorf("config type is invalid")
	}
	c, err := conf.doCheck()
	if err != nil {
		return err
	}
	k.lock.Lock()
	defer k.lock.Unlock()
	// 关闭旧的生产者
	if k.producer != nil {
		k.close()
	}
	k.producer, err = sarama.NewAsyncProducer(c.Address, c.Conf)
	if err != nil {
		return err
	}
	k.conf = c
	k.input = k.producer.Input()
	k.level = level
	k.formatType = conf.FormatType
	go k.sendMsg()
	return nil

}

func (k *Kafka) Output(infos []model.Info) error {
	k.lock.RLock()
	defer k.lock.RUnlock()
	if k.producer == nil || k.input == nil {
		return nil
	}
	if !k.enable {
		return nil
	}
	body, err := format.Format(k.formatType, infos)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: k.conf.Topic,
		Value: sarama.ByteEncoder(body),
	}
	if k.conf.PartitionType == "manual" {
		msg.Partition = k.conf.Partition
	}
	if k.conf.PartitionType == "hash" {
		msg.Key = sarama.StringEncoder(k.conf.PartitionKey)
	}
	log.Println("kafka 告警开始")
	k.input <- msg
	log.Println("kafka 告警结束")
	return nil
}

func (k *Kafka) Finish() error {
	k.lock.Lock()
	defer k.lock.Unlock()
	k.close()
	k.producer = nil
	k.input = nil
	return nil
}

func (k *Kafka) sendMsg() {
	if k.enable {
		return
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	k.cancel = cancelFunc
	// 初始化消息通道
	if k.wg == nil {
		k.wg = &sync.WaitGroup{}
	}
	k.enable = true
	// 监听关闭处理
	k.wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			// 读完
			for e := range k.producer.Errors() {
				log.Printf("kafka error:%s", e.Error())
			}
			k.wg.Done()
			return
		case err := <-k.producer.Errors():
			if err != nil {
				log.Printf("kafka error:%s", err.Error())
			}
		}
	}
}
func (k *Kafka) close() {
	if !k.enable {
		return
	}
	isClose := false
	k.producer.AsyncClose()
	if k.cancel != nil {
		isClose = true
		k.cancel()
		k.cancel = nil
	}
	if isClose {
		// 等待消息都读完
		k.wg.Wait()
	}
	k.producer = nil
	k.enable = false
}
