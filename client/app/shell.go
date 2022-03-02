package app

import (
	"github.com/gin-gonic/gin"
)

// 脚本执行
func (h *Http) Shell(c *gin.Context) {
	route := c.Param("R")              //读取shell路由配置
	if _, ok := h.Router[route]; !ok { //判断路由是否在配置文件中
		c.JSON(404, gin.H{
			"error": "no route",
		})
	} else { // 正常执行相应的shell
		c.JSON(200, gin.H{
			"route": route,
		})
	}
}
