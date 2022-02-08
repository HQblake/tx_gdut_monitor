package http

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"

// IAgent 页面管理Agent的相关接口
type IAgent interface {
	GetAgents() []*model.AgentInfo
}
