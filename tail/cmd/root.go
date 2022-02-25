package cmd

import (
	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "tail [command]",
	Short: "Gumiho tail",
	Long:  "Gumiho tail is hooks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
