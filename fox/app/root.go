package app

import "github.com/gin-gonic/gin"

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

	// 首页
	r.GET("/", h.home)

	return r
}
