package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BotToken string `yaml:"bot-token"`
	AdminID  int    `yaml:"admin-id"`
	Database struct {
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// GetConfig - функция для чтения конфигурации
func GetConfig(path string) (Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
