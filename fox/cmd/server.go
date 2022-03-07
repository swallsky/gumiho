package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"fox.com/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Server Fox

// 初始化
func init() {
	Server := Fox{}
	Server.initConfig()            //初始化配置
	Server.initResources()         //初始化资源
	Server.initCmd()               //初始化命令行
	Server.defStart()              // 定义命令开始
	rootCmd.AddCommand(Server.Cmd) // 添加服务根命令
	// app.Http 初始化
	Server.Http = app.Http{
		Host:       Server.Host,
		Port:       Server.Port,
		LogFile:    Server.LogFile,
		RootDir:    Server.RootDir,
		RuntimeDir: Server.RuntimeDir,
		StaticDir:  Server.StaticDir, //静态页面
	}
}

// 基类
type Fox struct {
	Host       string
	Port       string
	LogFile    string
	RootDir    string //项目根目录
	RuntimeDir string
	StaticDir  string //前台静态网面
	Cmd        *cobra.Command
	Http       app.Http
}

// 初始化配置
func (t *Fox) initConfig() {
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.SetConfigName("conf")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	t.LogFile = conf.GetString("server.logfile")       //监听的日志文件
	t.Host = conf.GetString("server.host")             //监听主机ip
	t.Port = conf.GetString("server.port")             //监听端口
	t.RuntimeDir = conf.GetString("server.runtimeDir") //运行时目录
	t.StaticDir = conf.GetString("server.StaticDir")   //前端网页
}

// 初始化资源
func (t *Fox) initResources() {
	// 创建目录
	if dir, err := os.Getwd(); err == nil {
		t.RootDir = dir
		t.RuntimeDir = dir + t.RuntimeDir
	}

	// fmt.Println(t.RuntimeDir)

	// 配置log日志文件
	t.LogFile = t.RuntimeDir + "/" + t.LogFile

	// 如果目录不存在，则创建
	if err := os.MkdirAll(t.RuntimeDir, 0777); err != nil {
		fmt.Println(err.Error())
	}
}

// 返回服务根命令
func (t *Fox) initCmd() {
	t.Cmd = &cobra.Command{
		Use:   "server", //命令标识
		Short: "Service management",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}

// 启动服务
func (t *Fox) defStart() {
	var upCmd = &cobra.Command{
		Use:   "up [flags]",
		Short: "server up [flags]",
		Run: func(cmd *cobra.Command, args []string) {
			daemon, _ := cmd.Flags().GetBool("daemon")
			if daemon { //是否启动daemon进程
				t.daemonPro()
			}
			t.Start() // 启动服务
		},
	}
	// 设置局部参数
	upCmd.Flags().BoolP("daemon", "d", false, "是否开启守护进程") // 是否开启守护进程
	t.Cmd.AddCommand(upCmd)
}

// 守护进程
func (t *Fox) daemonPro() {
	pid := syscall.Getppid()
	if pid == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		syscall.Umask(0)
		return
	}
	fmt.Println("Gumiho tail daemon start!")
	fp, err := os.OpenFile(t.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()
	c := exec.Command(os.Args[0], os.Args[1:]...)
	c.Stdout = fp
	c.Stderr = fp
	c.Stdin = nil
	c.SysProcAttr = &syscall.SysProcAttr{Setsid: true} //TODO TEST

	if err := c.Start(); err != nil {
		panic(err)
	}
	//将当前进程id写入文件中
	err = ioutil.WriteFile(t.RuntimeDir+"/.pid", []byte(fmt.Sprint(c.Process.Pid)), 0755)
	if err != nil {
		panic(err)
	}
	_, _ = fp.WriteString(fmt.Sprintf("[PID] %d Start At %s\n", c.Process.Pid, time.Now().Format("2006-01-02 15:04:05")))
	os.Exit(0)
}

// 启动服务
func (t *Fox) Start() {
	//路由初始化
	router := t.Http.InitRouter()
	// 服务初始化
	server := &http.Server{
		Addr:         ":" + t.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second, //读取超时时间
		WriteTimeout: 10 * time.Second, //写超时时间
	}

	go func() {
		// 连接服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 平滑重启
	listenSignal(server)
}

// 监听退出信号
func listenSignal(httpSrv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("Shutdown Server ...")

	//5秒后安全的退出程序
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shundown:", err)
	}
	log.Println("Server exiting")
}

// 关闭服务
func (t *Fox) Stop() {
	pid, err := ioutil.ReadFile(t.RuntimeDir + "/.pid")
	if err != nil {
		panic(err)
	}
	ppids := string(pid)             //将byte类型转换为string
	ppid, err := strconv.Atoi(ppids) //将string类型转换为int
	if err != nil {
		panic(err)
	}
	err = syscall.Kill(ppid, syscall.SIGINT) //关闭服务ctrl+c 2
	if err != nil {
		log.Println("Server logout fail!")
	} else {
		log.Println("Server logout successful!")
	}
}
