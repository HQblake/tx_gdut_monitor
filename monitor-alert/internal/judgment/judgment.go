package judgment

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
)

// JudgmentService 告警系统——判定服务模块
type JudgmentService struct{}

// 为告警系统的其他服务提供调用 JudgmentService 服务相应功能的接口

// Check 方法用于判定接入服务接收的agent上报数据是否正常
func (js *JudgmentService) Check(agent *model.AgentReport) error {
	return nil
}

// Update 方法接收管理服务发来的agent指标判定规则，更新服务内部缓存
func (js *JudgmentService) Update(rule *model.AgentRule) error {
	return nil
}
