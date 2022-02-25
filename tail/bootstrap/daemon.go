package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// 开启守护进程
func DaemonStart() {
	pid := syscall.Getppid()
	if pid == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		syscall.Umask(0)
		return
	}
	fmt.Println("Gumiho tail daemon start!")
	fp, err := os.OpenFile(Config.logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
	err = ioutil.WriteFile(Config.runtimeDir+"/.pid", []byte(fmt.Sprint(c.Process.Pid)), 0755)
	if err != nil {
		panic(err)
	}
	_, _ = fp.WriteString(fmt.Sprintf("[PID] %d Start At %s\n", c.Process.Pid, time.Now().Format("2006-01-02 15:04:05")))
	os.Exit(0)
}
