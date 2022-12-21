package config

import (
	"os"
	"gopkg.in/yaml.v3"
	"github.com/Dispanel/utils"
)

type Config struct {
	Port string `yaml:"port"`
	Authorization string `yaml:"authorization"`
	Logs struct {
		Discord bool `yaml:"discord"`
		channel string `yaml:"channel"`
		Uptime string `yaml:"uptimechannel"`
	} `yaml:"logs"`
}

func LoadConfig() Config {
	var config Config
	configFile, err := os.Open("../config.yaml")
	if err != nil {
		utils.ErrorHandling(err)
	}
	defer configFile.Close()
	yamlParser := yaml.NewDecoder(configFile)
	err = yamlParser.Decode(&config)
	if err != nil {
		utils.ErrorHandling(err)
	}
	return config
}