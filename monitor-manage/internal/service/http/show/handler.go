package show

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	// managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/judgment/gen"

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
			"msg":  "获取警告列表有误" + err.Error(),
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
	buff := bytes.NewBuffer([]byte{})
	buff.ReadFrom(c.Request.Body)
	hinfo := model.HistoryInfo{}
	err := json.Unmarshal(buff.Bytes(), &hinfo)
	if err != nil {
		log.Printf("json.Unmarshal();err: %v", err)
		return
	}

	log.Println(hinfo)
	format := "2006-01-02 15:04:05"
	start, err := time.ParseInLocation(format, hinfo.Start, time.Local)
	if err != nil {
		log.Printf("time.Parse();err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": "040004",
			"msg":  "初始时间和终止时间获取有误",
			"data": nil,
		})
		return
	}

	end, err := time.ParseInLocation(format, hinfo.End, time.Local)
	if err != nil {
		log.Printf("time.Parse();err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": "040004",
			"msg":  "初始时间和终止时间获取有误",
			"data": nil,
		})
		return
	}

	res, err := h.service.GetWarnInfoWithParams(hinfo, start, end)
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
	buff := bytes.NewBuffer([]byte{})
	buff.ReadFrom(c.Request.Body)
	metInfo := model.MetricsReq{}
	err := json.Unmarshal(buff.Bytes(), &metInfo)
	if err != nil {
		log.Printf("json.Unmarshal();err: %v", err)
		return
	}
	log.Println(metInfo)

	log.Printf(metInfo.IP, metInfo.Local, metInfo.MetricName, metInfo.Period, metInfo.Begin, metInfo.End, metInfo.Limit)

	format := "2006-01-02 15:04:05"
	newBegin := metInfo.Begin.Format(format)
	newEnd := metInfo.End.Format(format)
	metInfo.Begin = newBegin
	metInfo.End = newEnd

	res, err := h.service.GetMetricsWithTime()
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
