package agent

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service agent.IAgent
}
func NewHandler(service agent.IAgent) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllAgent(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) GetAgentInfo(c *gin.Context) {
	panic("implement me")
}

