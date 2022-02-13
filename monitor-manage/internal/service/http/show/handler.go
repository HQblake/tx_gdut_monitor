package show

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetWarnInfo(c *gin.Context) {
	res, err := h.service.GetWarnInfo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040001",
			"msg":  "获取警告列表有误",
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

func (h *Handler) GetWarnInfoWithId(c *gin.Context) {
	Id := c.Param("id")
	if Id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "040002",
			"msg":  "Id获取有误",
			"data": nil,
		})
		return
	}
	NewId, _ := strconv.ParseInt(Id, 10, 32)
	res, err := h.service.GetWarnInfoWithId(int32(NewId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040003",
			"msg":  "ID获取告警信息有误",
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

func (h *Handler) GetWarnInfoWithParams(c *gin.Context) {
	// ip string, local string, level int32, begin int64, end int64
	begin := c.Query("begin")
	end := c.Query("end")
	if begin == "" || end == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "040004",
			"msg":  "初始时间和终止时间获取有误",
			"data": nil,
		})
		return
	}
	ip := c.Query("ip")
	local := c.Query("local")
	metric := c.Query("metric")
	level := c.Query("level")
	var newLevel int64
	if level != "" {
		newLevel, _ = strconv.ParseInt(level, 10, 32)
	} else {
		newLevel = 0
	}
	newBegin, _ := strconv.ParseInt(begin, 10, 64)
	newEnd, _ := strconv.ParseInt(end, 10, 64)
	res, err := h.service.GetWarnInfoWithParams(ip, local, metric, int32(newLevel), newBegin, newEnd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040005",
			"msg":  "根据参数查询告警列表出错",
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

func (h *Handler) GetMetricsWithTime(c *gin.Context) {
	ip := c.Query("ip")
	local := c.Query("local")
	metric := c.Query("metric")
	begin := c.Query("begin")
	limit := c.Query("limit")
	if ip == "" || local == "" || metric == "" || begin == "" || limit == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "040006",
			"msg":  "获取主机指标参数有误",
			"data": nil,
		})
		return
	}
	newBegin, _ := strconv.ParseInt(begin, 10, 64)
	newLimit, _ := strconv.ParseInt(limit, 10, 64)
	res, err := h.service.GetMetricsWithTime(ip, local, metric, newBegin, int32(newLimit))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040007",
			"msg":  "获取主机指标信息有误",
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

func (h *Handler) DelWarnInfo(c *gin.Context) {

	id := c.Param("id")
	delId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040008",
			"msg":  "参数id信息有误，请重试" + err.Error(),
			"data": nil,
		})
		return
	}
	err = h.service.DelWarnInfo(int32(delId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "040008",
			"msg":  "删除告警信息出错" + err.Error(),
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
