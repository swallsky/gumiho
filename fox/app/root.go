package app

import (
	"github.com/gin-gonic/gin"
)

type Http struct {
	Host       string
	Port       string
	LogFile    string
	RootDir    string //项目根目录
	RuntimeDir string
	Router     map[string]string // 路由配置
}

func (h *Http) InitRouter() *gin.Engine {
	// 设置模式
	// gin.SetMode(gin.DebugMode) //默认模式
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	r := gin.Default()

	// 设置静态目录
	// r.Use(static.Serve("/", static.LocalFile(h.StaticDir, true)))
	r.GET("/", h.home)     //首页
	r.GET("/pong", h.pong) //测试
	// api
	// api := r.Group("/api")
	// {
	// 	api.GET("/", h.home)     // api首页
	// 	api.GET("/pong", h.pong) // api pong test
	// }

	return r
}
