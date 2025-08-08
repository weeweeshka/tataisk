package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	StoragePath string `yaml:"storage_path"`
}

type GRPC struct {
	Port    string `yaml:"port"`
	Timeout string `yaml:"timeout"`
}

func MustLoadConfig() *Config {
	configPath := "config/local.yaml"

	if _, err := os.Stat("local.yaml"); os.IsNotExist(err) {
		panic("local.yaml doesn't exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
