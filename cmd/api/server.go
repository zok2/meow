package api

import (
	"github.com/spf13/cobra"
	"github.com/zok2/meow/core"
	"github.com/zok2/meow/global"
	"github.com/zok2/meow/initialize"
	"go.uber.org/zap"
	"log"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "meow server -c config/settings.yml",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

}

func setup() {
	//gorm 连接数据库
	global.GVA_VP = core.Viper(configYml) // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.GormMysql()
	usageStr := `starting api server...`
	log.Println(usageStr)
}

func run() error {
	core.RunWindowsServer()
	return nil
}
