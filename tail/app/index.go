package app

import (
	"github.com/gin-gonic/gin"
)

// 默认首页
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "tail hooks",
	})
}
