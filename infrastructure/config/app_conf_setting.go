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
	profile := os.Getenv("profile")
	var filename string
	if profile == "" {
		filename = "./conf/app.yaml"
	} else {
		filename = fmt.Sprintf("./conf/app_%s.yaml", profile)
	}
	yamlFile, err := os.ReadFile(filename)
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
