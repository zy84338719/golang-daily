package config

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"spider/utils"
)

var Conf *viper.Viper

func init() {
	Conf = viper.New()

	Conf.SetConfigName("config")
	path:= utils.ExecPath
	Conf.AddConfigPath(path)

	Conf.SetConfigType("yaml")
	if err := Conf.ReadInConfig(); err != nil {
		log.Error("配置初始化异常，",err)
		panic(err)
	}
}