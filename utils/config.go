package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ServerConfig   ServerConfig   `yaml:"server"`
	RabbitMQConfig RabbitMQConfig `yaml:"rabbit_mq"`
}

//ServerConfig : ServerConfig
type ServerConfig struct {
	Port string `yaml:"port"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//InitServiceConfig : Init service config from local file path
func InitServiceConfig() (*AppConfig, error) {
	config, err := parseConfigFromPath("config.yaml")
	if err != nil {
		return nil, err
	}
	return config, nil
}

func parseConfigFromPath(configPath string) (*AppConfig, error) {
	var config = new(AppConfig)
	fmt.Printf("get service config from config path: %v\n", configPath)
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return config, nil
}
