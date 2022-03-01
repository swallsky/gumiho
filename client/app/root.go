package app

import "github.com/gin-gonic/gin"

type Http struct {
	Host       string
	Port       string
	LogFile    string
	RuntimeDir string
}

func (h *Http) InitRouter() *gin.Engine {
	// 设置模式
	// gin.SetMode(gin.DebugMode) //默认模式
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	r := gin.Default()

	// 首页
	r.GET("/", h.Home)

	// 执行远程脚本
	// r.GET("/shell/:R", app.Shell)

	return r
}
