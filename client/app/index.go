package app

import (
	"github.com/gin-gonic/gin"
)

func (h *Http) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "tail hooks",
	})
}
