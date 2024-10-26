package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// This struct reflects the yaml file structure
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Stores struct {
		MySql struct {
			ConnectionString string `yaml:"connection_string"`
		}
	}
}

var Configs Config

func LoadConfig() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &Configs)
	if err != nil {
		panic(err)
	}
}