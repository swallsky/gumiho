package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"tail.com/conf"
)

// 配置的路由
var Router map[string]string

// 初始化
func init() {
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.SetConfigName("conf")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	Router = make(map[string]string, 10)
	Router = conf.GetStringMapString("router")
}

// 执行相应的shell
func shellBin(name string, shell string) {
	path := conf.Server.runtimeDir
	// shellFile := path.Join("")
}

func Shell(c *gin.Context) {
	route := c.Param("R")            //读取shell路由配置
	if _, ok := Router[route]; !ok { //判断路由是否在配置文件中
		c.JSON(404, gin.H{
			"error": "no route",
		})
	} else { // 正常执行相应的shell
		c.JSON(200, gin.H{
			"route": route,
		})
	}
}
