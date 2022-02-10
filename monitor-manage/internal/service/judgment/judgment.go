package judgment

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
)

// IJudgment 服务接口化，对外提供Update方法，内部实际调用Grpc，尽可能解耦
type IJudgment interface {
	GetConfigs(ip string, local string) ([]model.JudgmentConfig,map[string]model.JudgmentConfig, error)
	// Update 更新判定服务配置，一方面更新数据库，一方面告知判定服务rpc
	Update(config model.JudgmentConfig) error
	// Del 删除指定id的配置
	Del(ip string, local string, id int32) error
}


