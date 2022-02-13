package setting

var cfg = new(config)

type config struct {
	AgentConfig   *agentConfig   `json:"agent" yaml:"agent"`
	ConnectConfig *connectConfig `json:"connect" yaml:"connect"`
}

type connectConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type agentConfig struct {
	Internal  int64    `json:"interval" yaml:"interval"`
	Metric    []string `json:"metric" yaml:"metric"`
	Dimension []string `json:"dimension" yaml:"dimension"`
	Location  string   `json:"location" yaml:"location"`
}
