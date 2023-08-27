package initialize

import (
	"fmt"
	"github.com/life-studied/douyin-simple/global"
	"github.com/spf13/viper"
	"os"
)

// Read  ReadConfig
func Read(configName string, configPath string, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

// Config InitConfig
func Config() {
	path, err := os.Getwd() //返回项目目录
	if err != nil {
		panic(err)
	}
	v := Read("config", path, "yaml") //在项目目录底下查找config.yaml配置文件
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
}
