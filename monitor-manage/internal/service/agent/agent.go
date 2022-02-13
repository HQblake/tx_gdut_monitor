package agent

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
)

type IAgent interface {
	GetAllAgentInfo() ([]model.AgentInfo, error)
	GetAllAgentSendInfo() ([]model.AgentSendInfo, error)
	GetAgentInfo(ip string, local string) (*model.AgentInfo, error)
}

