package api

import (
	"github.com/spf13/cobra"
	"go-application/tools/config"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API Server",
		Example:      "go-application server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/setting.yml", "Start server with provided configuration file")
}

func setup() {
	//1. 读取配置
	config.Setup(configYml)

	//2. 启动数据库
}
