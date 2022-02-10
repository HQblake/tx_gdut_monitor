package show

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service show.IShow
}
func NewHandler(service show.IShow) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMetricsInOneDay(c *gin.Context) {
	panic("implement me")
}


