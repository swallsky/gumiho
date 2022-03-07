package app

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Http struct {
	Host       string
	Port       string
	LogFile    string
	RootDir    string //项目根目录
	RuntimeDir string
	StaticDir  string            //静态页面路径
	Router     map[string]string // 路由配置
}

func (h *Http) InitRouter() *gin.Engine {
	// 设置模式
	// gin.SetMode(gin.DebugMode) //默认模式
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	r := gin.Default()

	// 设置静态目录
	r.Use(static.Serve("/", static.LocalFile(h.StaticDir, true)))

	// api
	api := r.Group("/api")
	{
		api.GET("/", h.home)     // api首页
		api.GET("/pong", h.pong) // api pong test
	}

	return r
}
