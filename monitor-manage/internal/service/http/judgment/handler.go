package judgment

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service judgment.IJudgment
	agent agent.IAgent
}

func NewHandler(service judgment.IJudgment, agent agent.IAgent) *Handler {
	return &Handler{
		service: service,
		agent: agent,
	}
}

func (h *Handler) GetAllRule(c *gin.Context) {
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
	agentInfo, err := h.agent.GetAgentInfo(ip, local)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "020003",
			"msg":  "获取agent信息出错" + err.Error(),
			"data": nil,
		})
		return
	}
	res, err := h.service.GetConfigsWithMetrics(ip, local, agentInfo.Metric)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "020001",
			"msg":  "获取agent rule信息出错" + err.Error(),
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

func (h *Handler) UpdateRule(c *gin.Context) {
	id := c.Param("id")
	updateId , err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数id信息有误，请重试"  + err.Error(),
			"data": nil,
		})
		return
	}
	method := c.PostForm("method")
	updateMethod, err := strconv.Atoi(method)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数method信息有误，请重试"  + err.Error(),
			"data": nil,
		})
		return
	}
	ip := c.Param("ip")
	local := c.Param("local")
	metric := c.PostForm("metric")
	period := c.PostForm("period")
	threshold := c.PostForm("threshold")
	if ip == "" || local == "" || metric == "" || period == "" || threshold == ""{
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	err = h.service.Update(int32(updateId), ip, local, metric, int32(updateMethod), period, threshold)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "020002",
			"msg":  "更新agent rule信息出错"  + err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
	})
	return
}

func (h *Handler) DelRule(c *gin.Context) {
	ip := c.Param("ip")
	local := c.Param("local")

	if ip == "" || local == ""{
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	id := c.Param("id")
	delId , err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数id信息有误，请重试"  + err.Error(),
			"data": nil,
		})
		return
	}
	err = h.service.Del(ip, local, int32(delId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "020003",
			"msg":  "删除agent rule信息出错"  + err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
	})
	return
}

