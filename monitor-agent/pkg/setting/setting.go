package setting

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func InitConfig(path string) error { //初始化读取配置文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, cfg)
}

func GetDefaultConfig() *config {
	return cfg
}

func GetAgentInterval() int64 { //获取Agent上报时间间隔
	return cfg.AgentConfig.Internal
}

func GetAgentMetrics() []string { //获取Agent需要上报的指标数组
	return cfg.AgentConfig.Metric
}

func GetAgentDimensions() []string { //获取Agent需要上报的维度数组
	return cfg.AgentConfig.Dimension
}

func GetConnectHost() string { //获取服务端的IP地址
	return cfg.ConnectConfig.Host
}

func GetConnectPort() string { //获取服务端grpc监听的端口号
	return cfg.ConnectConfig.Port
}

func GetConnectLocation() string { //获取Agent端的区域信息
	return cfg.AgentConfig.Location
}
