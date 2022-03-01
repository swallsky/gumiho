package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// 服务相关的配置
var Server struct {
	LogFile    string
	Host       string
	Port       string
	RuntimeDir string
}

// 初始化
func init() {
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.SetConfigName("conf")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	Server.LogFile = conf.GetString("server.logfile")       //监听的日志文件
	Server.Host = conf.GetString("server.host")             //监听主机ip
	Server.Port = conf.GetString("server.port")             //监听端口
	Server.RuntimeDir = conf.GetString("server.runtimeDir") //运行时目录
}

// 创建目录
func init() {
	if dir, err := os.Getwd(); err == nil {
		Server.RuntimeDir = dir + Server.RuntimeDir
	}

	// 配置log日志文件
	Server.LogFile = Server.RuntimeDir + "/" + Server.LogFile

	// 如果目录不存在，则创建
	if err := os.MkdirAll(Server.RuntimeDir, 0777); err != nil {
		fmt.Println(err.Error())
	}
}
