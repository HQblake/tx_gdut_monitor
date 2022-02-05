package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

type Config struct {
	Topic         string   `json:"topic"`
	Address       []string `json:"address"`
	Version       string   `json:"version"`
	PartitionType string   `json:"partition_type"`
	Partition     int32    `json:"partition"`
	PartitionKey  string   `json:"partition_key"`
	FormatType    string   `json:"format_type"`
}
type ProducerConfig struct {
	Address       []string       `json:"address"`
	Topic         string         `json:"topic"`
	Partition     int32          `json:"partition"`
	PartitionKey  string         `json:"partition_key"`
	PartitionType string         `json:"partition_type"`
	Conf          *sarama.Config `json:"conf"`
	FormatType    string         `json:"format_type"`
}

func (c *Config) doCheck() (*ProducerConfig, error) {
	c.FormatType = strings.ToLower(c.FormatType)
	if c.FormatType == "" {
		c.FormatType = "line"
	}
	if c.Topic == "" {
		return nil, fmt.Errorf("topic can not be null")
	}
	if len(c.Address) == 0 {
		return nil, fmt.Errorf("address can not be null")
	}
	p := &ProducerConfig{}
	p.Address = c.Address
	p.Topic = c.Topic
	s := sarama.NewConfig()
	if c.Version != "" {
		v, err := sarama.ParseKafkaVersion(c.Version)
		if err != nil {
			return nil, err
		}
		s.Version = v
	}
	p.PartitionType = c.PartitionType
	switch c.PartitionType {
	case "robin":
		s.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	case "hash":
		// 通过hash获取
		if c.PartitionKey == "" {
			// key为空则还是用随机
			s.Producer.Partitioner = sarama.NewRandomPartitioner
			p.PartitionType = "random"
		} else {
			s.Producer.Partitioner = sarama.NewHashPartitioner
			p.PartitionKey = c.PartitionKey
		}
	case "manual":
		// 手动指定分区
		s.Producer.Partitioner = sarama.NewManualPartitioner
		// 默认为0
		p.Partition = c.Partition
	default:
		s.Producer.Partitioner = sarama.NewRandomPartitioner
		p.PartitionType = "random"
	}
	// 只监听错误
	s.Producer.Return.Errors = true
	s.Producer.Return.Successes = false
	s.Producer.RequiredAcks = sarama.WaitForLocal
	p.Conf = s
	return p, nil
}
