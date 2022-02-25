package bootstrap

import "github.com/spf13/viper"

// 服务相关的配置
var Config struct {
	logfile string
	host    string
	port    string
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
	Config.logfile = conf.GetString("server.logfile") //监听的日志文件
	Config.host = conf.GetString("server.host")       //监听主机ip
	Config.port = conf.GetString("server.port")       //监听端口
}
