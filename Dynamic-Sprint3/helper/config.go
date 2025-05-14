package helper

import (
	"dynamic/model"
	"encoding/json"
	"fmt"
	"os"
)

var configFilePath string
var config *model.Config

func LoadConfig() *model.Config {
	var configModel *model.Config
	configFileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
		return &model.Config{}
	}
	err = json.Unmarshal(configFileBytes, &configModel)
	if err != nil {
		fmt.Println(err)
		return &model.Config{}
	}
	return configModel
}

func GetConfig() *model.Config {
	if config == nil {
		config = LoadConfig()
	}
	return config
}

func DefineConfigPath(configFile string) {
	configFilePath = configFile
}
