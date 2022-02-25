package cmd

import (
	"github.com/spf13/cobra"
	"tail.com/bootstrap"
)

// 当前命令
var serverCmd = &cobra.Command{
	Use:   "server", //命令标识
	Short: "Service management",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// 启动服务
var upCmd = &cobra.Command{
	Use:   "up [flags]",
	Short: "server up [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		if daemon { //是否启动daemon进程
			bootstrap.DaemonStart()
		}
		bootstrap.ServerStart()
	},
}

// 停止服务
var downCmd = &cobra.Command{
	Use:   "stop",
	Short: "Server stop ...",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.ServerStop()
	},
}

// 初始化
func init() {
	// 变量参数定义 持久性flags 全局
	// serverCmd.PersistentFlags().BoolVar(&daemon, "d", false, "true: daemon start,default is false")

	// 设置局部参数
	upCmd.Flags().BoolP("daemon", "d", false, "是否开启守护进程") // 是否开启守护进程
	serverCmd.AddCommand(upCmd)                           //启动服务
	serverCmd.AddCommand(downCmd)                         // 停止服务
	rootCmd.AddCommand(serverCmd)
}
