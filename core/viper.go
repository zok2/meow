package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zok2/meow/core/internal"
	"github.com/zok2/meow/global"
	"os"
)

func Viper(config string) *viper.Viper {
	if config == "" {
		if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
			switch gin.Mode() {
			case gin.DebugMode:
				config = internal.ConfigDefaultFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
			case gin.ReleaseMode:
				config = internal.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
			case gin.TestMode:
				config = internal.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
			}
		} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
			config = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
