package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Config   = viper.New()
	confPath = "conf/"
)

// to do
// alone instance config

func InitConfig() {
	var err error

	Config.AddConfigPath(confPath)
	Config.SetConfigName("es")
	err = Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error es file config: %s \n", err.Error()))
	}

	Config.SetConfigName("base")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file base: %s \n", err.Error()))
	}

	Config.SetConfigName("log")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file log: %s \n", err.Error()))
	}

	Config.SetConfigName("redis")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file redis: %s \n", err.Error()))
	}

	Config.SetConfigName("database")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file database: %s \n", err.Error()))
	}


	Log.Info("初始化日志配置")
}