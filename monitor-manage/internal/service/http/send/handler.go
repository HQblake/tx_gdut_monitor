package send

import (
	"net/http"
	"strconv"

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
func (h *Handler) GetSendConfigs(c *gin.Context) {
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
	res, err := h.service.GetConfigs(ip, local)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "030001",
			"msg":  "获取agent发送配置信息出错" + err.Error(),
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

func (h *Handler) AddSendConfig(c *gin.Context) {
	sendType := c.PostForm("sendType")
	addSendType, err := strconv.Atoi(sendType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数sendType信息有误，请重试",
			"data": nil,
		})
		return
	}
	level := c.PostForm("level")
	addLevel, err := strconv.Atoi(level)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数level信息有误，请重试" + err.Error(),
			"data": nil,
		})
		return
	}
	ip := c.Param("ip")
	local := c.Param("local")
	config := c.PostForm("config")
	if ip == "" || local == "" || config == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	err = h.service.AddConfig(ip, local, int32(addSendType), int32(addLevel), config)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "030002",
			"msg":  "新增agent发送配置出错" + err.Error(),
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

func (h *Handler) UpdateSendConfig(c *gin.Context) {
	id := c.Param("id")
	updateId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数id信息有误，请重试",
			"data": nil,
		})
		return
	}
	sendType := c.PostForm("sendType")
	updateSendType, err := strconv.Atoi(sendType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数sendType信息有误，请重试",
			"data": nil,
		})
		return
	}
	level := c.PostForm("level")
	updateLevel, err := strconv.Atoi(level)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数level信息有误，请重试",
			"data": nil,
		})
		return
	}
	ip := c.Param("ip")
	local := c.Param("local")
	config := c.PostForm("config")
	if ip == "" || local == "" || config == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数信息有误，请重试",
			"data": nil,
		})
		return
	}
	err = h.service.Update(int32(updateId), ip, local, int32(updateSendType), int32(updateLevel), config)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "030003",
			"msg":  "更新agent发送配置出错" + err.Error(),
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

func (h *Handler) DelSendConfig(c *gin.Context) {
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
	id := c.Param("id")
	delId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "000001",
			"msg":  "参数id信息有误，请重试" + err.Error(),
			"data": nil,
		})
		return
	}
	err = h.service.Del(ip, local, int32(delId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "030004",
			"msg":  "删除agent发送配置出错" + err.Error(),
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
