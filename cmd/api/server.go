package api

import (
	"fmt"
	"go-application/common/database"
	"go-application/tools/config"

	"github.com/spf13/cobra"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
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
	fmt.Println(config.DatabaseConfig.Driver)
	database.Setup(config.DatabaseConfig.Driver)
}

func run() error {
	fmt.Println("asdf")
	return fmt.Errorf("aa")
}
