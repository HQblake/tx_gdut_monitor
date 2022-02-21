package setting


type AlertConfig struct {
	Convergence int `json:"convergence" yaml:"convergence"` // 告警收敛方式，0为即时告警，1为聚合收敛，2为滚动收敛(未完善)，默认为1
	Interval int `json:"interval" yaml:"interval"` // 告警周期，单位为min，收敛方式为1或2时有效，默认为1min
}