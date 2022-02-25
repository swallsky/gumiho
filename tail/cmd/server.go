package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "server start [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		fmt.Println(daemon)
		fmt.Println("#v", args)
	},
}

// 初始化
func init() {
	// 变量参数定义 定义全局
	// serverCmd.PersistentFlags().BoolVar(&daemon, "d", false, "true: daemon start,default is false")

	// 设置局部参数
	startCmd.Flags().BoolP("daemon", "d", false, "是否开启守护进程")
	serverCmd.AddCommand(startCmd)
	rootCmd.AddCommand(serverCmd)
}
