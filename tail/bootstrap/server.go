package bootstrap

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// 启动服务
// port 启动端口号
func ServerStart() {
	// host := Config.host
	port := Config.port
	//路由初始化
	router := InitRouter()
	// 服务初始化
	server := &http.Server{
		Addr:         ":" + port,
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
func ServerStop() {
	// host := Config.host
	// port := Config.port
	pid, err := ioutil.ReadFile("./runtime/.pid")
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
