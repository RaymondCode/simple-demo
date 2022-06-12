package dto

import (
	"io/ioutil"

	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env struct {
		IsDebug bool `yaml:"debug"`
	} `yaml:"env"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	MySQL struct {
		Local struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"local"`
		Default struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"default"`
	} `yaml:"mysql"`
	Redis struct {
		Local struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Password string `yaml:"password"`
		} `yaml:"local"`
		Default struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Password string `yaml:"password"`
		} `yaml:"default"`
		Databases map[string]int `yaml:"databases"`
	} `yaml:"redis"`
	Log struct {
		Level      string `yaml:"level"`
		Filename   string `yaml:"filename"`
		MaxSize    int    `yaml:"max_size"`
		MaxAge     int    `yaml:"max_age"`
		MaxBackups int    `yaml:"max_backups"`
	} `yaml:"log"`
}

var Conf *Config

func GetConfig() *Config {
	return Conf
}

func InitConfig() error {
	var config Config

	configFilePath := util.GetConfigPath()
	file, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, &config); err != nil {
		return err
	}

	Conf = &config
	return nil
}
