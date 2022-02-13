package http

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	agent2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http/agent"
	judgment2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http/judgment"
	send2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http/send"
	show2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http/show"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	IAgentHandler
	IJudgmentHandler
	ISendHandler
	IShowHandler
}

func NewHandler(agent agent.IAgent, judgment judgment.IJudgment, send send.ISend, show show.IShow) *Handler {
	return &Handler{
		IAgentHandler: agent2.NewHandler(agent),
		IJudgmentHandler: judgment2.NewHandler(judgment, agent),
		ISendHandler: send2.NewHandler(send),
		IShowHandler: show2.NewHandler(show),
	}
}

// IAgentHandler agent信息的管理http接口
type IAgentHandler interface {
	// GetAllAgent 获取所有的存活agent
	GetAllAgent(c *gin.Context)
	// GetAllAgentSendInfo 获取所有的存活agent
	GetAllAgentSendInfo(c *gin.Context)
	// GetAgentInfo 获取指定Agent的信息
	GetAgentInfo(c *gin.Context)
}


// IJudgmentHandler 判定服务的管理http接口
type IJudgmentHandler interface {
	// GetAllRule 获取指定agent的规则列表
	GetAllRule(c *gin.Context)
	// UpdateRule 更新指定id的判定规则
	UpdateRule(c *gin.Context)
	// DelRule 删除规则
	DelRule(c *gin.Context)
}

// ISendHandler 发送服务的管理http接口
type ISendHandler interface {
	// GetConfigs 获取指定agent的发送配置列表
	GetSendConfigs(c *gin.Context)
	// AddSendConfig  新增配置
	AddSendConfig(c *gin.Context)
	// UpdateSendConfig  更新发送服务配置
	UpdateSendConfig(c *gin.Context)
	// DelSendConfig 删除指定id的配置
	DelSendConfig(c *gin.Context)
}

// IShowHandler 视图显示的http接口，自青补充
type IShowHandler interface {
	GetMetricsInOneDay(c *gin.Context)
}

