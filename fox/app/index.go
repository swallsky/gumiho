package app

import (
	"github.com/gin-gonic/gin"
)

// 首页
func (h *Http) home(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "fox hooks",
	})
}

// 测试
func (h *Http) pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
