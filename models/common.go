package models

import (
	"os"

	"gopkg.in/yaml.v3"
)

type MySQLConfig struct {
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
}

type GlobalConfig struct {
	MySQLConf MySQLConfig `yaml:"mysql"`
}

func (c *GlobalConfig) getConf() *GlobalConfig {
	yamlFile, err := os.ReadFile("resources/application.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}

func InitProject() (err error) {
	var c GlobalConfig
	conf := c.getConf()
	err = InitMySql(conf.MySQLConf)
	if err != nil {
		panic(err)
	}
	return nil
}

func Close() {
	CloseMySQL()
}
