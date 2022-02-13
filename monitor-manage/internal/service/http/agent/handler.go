package agent

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service agent.IAgent
}
func NewHandler(service agent.IAgent) *Handler {
	return &Handler{
		service: service,
	}
}

// GetAllAgent Get
func (h *Handler) GetAllAgent(c *gin.Context) {
	res, err := h.service.GetAllAgentInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "010001",
			"msg":  "获取agent列表出错" + err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
		"data": res,
	})
	return
}

func (h *Handler) GetAllAgentSendInfo(c *gin.Context) {
	res, err := h.service.GetAllAgentSendInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "010001",
			"msg":  "获取agent列表出错" + err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
		"data": res,
	})
	return
}

// GetAgentInfo Get,query:ip,local
func (h *Handler) GetAgentInfo(c *gin.Context) {
	ip := c.Param("ip")
	local := c.Param("local")
	if ip == "" || local == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	res, err := h.service.GetAgentInfo(ip, local)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "010002",
			"msg":  "获取agent信息出错",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
		"data": res,
	})
	return
}

