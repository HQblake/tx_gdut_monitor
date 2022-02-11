package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
)

type ISend interface {
	// GetAllConfigs 获取所有配置
	GetAllConfigs() ([]model.SendConfig, error)
	// GetConfigs 获取配置列表
	GetConfigs(ip string, local string) ([]model.SendConfig, error)
	// AddConfig 新增配置
	AddConfig(Ip string, Local string, SendType int32, Level int32, Config string) error
	// Update 更新发送服务配置
	Update(id int32, IP string, Local string, SendType int32, Level int32, Config string) error
	// Del 删除指定id的配置
	Del(ip string, local string, id int32) error
	// Init 初始化
	Init()
}

