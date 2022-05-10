package config

import (
	"fmt"
	G "github.com/NoCLin/douyin-backend-go/config/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("dev")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Configuration file changed")
	})

	if err := viper.Unmarshal(&G.Config); err != nil {
		panic(fmt.Errorf("fatal error while decode config file: %v", err))
	}

	fmt.Println(G.Config)
	err := initGorm()
	if err != nil {
		return err
	}
	return nil
}
