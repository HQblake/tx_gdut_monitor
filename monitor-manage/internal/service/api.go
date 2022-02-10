package service

import (
	"github.com/gin-gonic/gin"
)

// ICheckHandler 判定服务的管理接口
type ICheckHandler interface {
	// ListAll 获取指定agent的规则列表
	ListAll(c *gin.Context)
	// Update 更新指定id的判定规则
	Update(c *gin.Context)
	// Del 删除规则
	Del(c *gin.Context)

}


