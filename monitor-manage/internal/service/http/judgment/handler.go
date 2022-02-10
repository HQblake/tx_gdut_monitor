package judgment

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service judgment.IJudgment
}

func NewHandler(service judgment.IJudgment) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllRule(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) UpdateRule(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) DelRule(c *gin.Context) {
	panic("implement me")
}

