package app

import (
	"github.com/gin-gonic/gin"
)

func (h *Http) home(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "fox hooks",
	})
}
