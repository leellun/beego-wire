package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"zhiqu/infrastructure/config/model"
)

var (
	AppConfig *model.Config
)

func InitConfig() {
	yamlFile, err := os.ReadFile("./conf/app.yaml")
	if err != nil {
		panic(err)
	}
	var _config *model.Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		panic(err)
	}
	AppConfig = _config
	fmt.Printf("config.app: %#v\n", _config)
}
