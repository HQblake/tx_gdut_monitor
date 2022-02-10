package send

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"

type ISend interface {
	AddConfig() error
	GetConfigs(ip string, local string) ([]model.SendConfig, error)
	// Update 更新发送服务配置
	Update(config model.SendConfig) error
	// Del 删除指定id的配置
	Del(ip string, local string, id int32) error

	Init(configs []model.SendConfig) error
}