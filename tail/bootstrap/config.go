package bootstrap

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// 服务相关的配置
var Config struct {
	logfile    string
	host       string
	port       string
	runtimeDir string
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
	Config.logfile = conf.GetString("server.logfile")       //监听的日志文件
	Config.host = conf.GetString("server.host")             //监听主机ip
	Config.port = conf.GetString("server.port")             //监听端口
	Config.runtimeDir = conf.GetString("server.runtimeDir") //运行时目录
}

// 创建目录
func init() {
	if dir, err := os.Getwd(); err == nil {
		Config.runtimeDir = dir + Config.runtimeDir
	}

	// 配置log日志文件
	Config.logfile = Config.runtimeDir + "/" + Config.logfile
	// fmt.Println(Config.runtimeDir)
	// fmt.Println(Config.logfile)
	// 如果目录不存在，则创建
	if err := os.MkdirAll(Config.runtimeDir, 0777); err != nil {
		fmt.Println(err.Error())
	}
}
