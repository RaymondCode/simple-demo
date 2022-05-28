package initialize

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
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
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	v := Read("config", path, "yaml")
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
}
