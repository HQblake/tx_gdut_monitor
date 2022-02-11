package show

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
	"github.com/gin-gonic/gin"
	"net/http"
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
	ip := c.Query("ip")
	local := c.Query("local")
	metric := c.Query("metric")
	if ip == "" || local == "" || metric == ""{
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	res := h.service.GetMetricsInOneDay(ip, local, metric)
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
		"data": res,
	})
	return
}


