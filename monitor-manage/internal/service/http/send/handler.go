package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service send.ISend
}
func NewHandler(service send.ISend) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) GetConfigs(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) AddSendConfig(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) UpdateSendConfig(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) DelSendConfig(c *gin.Context) {
	panic("implement me")
}

