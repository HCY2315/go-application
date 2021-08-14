package api

import (
	"fmt"
	"go-application/app/admin/router"
	"go-application/common/database"
	"go-application/common/global"
	"go-application/common/log"
	"go-application/pkg/cache"
	mycasbin "go-application/pkg/casbin"
	"go-application/tools"
	"go-application/tools/config"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
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

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/setting.yml", "Start server with provided configuration file")
}

func setup() {
	//1. 读取配置
	config.Setup(configYml)

	//2. 启动数据库
	database.Setup(config.DatabaseConfig.Driver)

	//3. 启动缓存服务器
	cache.SetUp()

	//4. 接口访问控制加载
	global.CasbinEnforcer = mycasbin.Setup(global.Eloquent, "sys", "casbin")

	usageStr := `starting api server !!!`
	log.Info("\n" + tools.Green(usageStr))
}

func run() error {
	if config.ApplicationConfig.Model == string(tools.ModelPord) {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := global.Cfg.GetEngine()
	if engine == nil {
		engine = gin.New()
	}

	if config.ApplicationConfig.Model == string(tools.ModelDev) {
		// 添加监控
		AppRouters = append(AppRouters, router.Monitor)
	}

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: global.Cfg.GetEngine(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("关闭服务")
	return fmt.Errorf("aa")
}
