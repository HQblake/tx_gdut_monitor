package send

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"

type ISend interface {
	// GetConfigs 获取配置列表
	GetConfigs(ip string, local string) ([]model.SendConfig, error)
	// AddConfig 新增配置
	AddConfig(Ip string, Local string, SendType int, Level int, Config string) error
	// Update 更新发送服务配置
	Update(config model.SendConfig) error
	// Del 删除指定id的配置
	Del(ip string, local string, id int32) error
	// Init 初始化
	Init(configs []model.SendConfig) error
}