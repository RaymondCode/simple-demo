package dto

import (
	"io/ioutil"

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
}

var Conf *Config

func GetConfig() *Config {
	return Conf
}

func InitConfig() error {
	var config Config

	file, err := ioutil.ReadFile("../config/config.yml")

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, &config); err != nil {
		return err
	}

	Conf = &config
	return nil
}
