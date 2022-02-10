package service

import (
	"github.com/gin-gonic/gin"
)

// IAgentHandler agent信息的管理http接口
type IAgentHandler interface {
	// GetAllAgent 获取所有的存活agent
	GetAllAgent(c *gin.Context)
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
	GetConfigs(c *gin.Context)
	// AddSendConfig  新增配置
	AddSendConfig(c *gin.Context)
	// UpdateSendConfig  更新发送服务配置
	UpdateSendConfig(c *gin.Context)
	// DelSendConfig 删除指定id的配置
	DelSendConfig(c *gin.Context)
}

// IShowHandler 视图显示的http接口，自青补充
type IShowHandler interface {

}

