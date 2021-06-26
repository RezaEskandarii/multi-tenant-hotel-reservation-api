package config

import (
	"gopkg.in/yaml.v2"
	"hotel-reservation/pkg/application_loger"
	"io/ioutil"
)

// Config application config struct
type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		SSLMode  string `yaml:"ssl_mode"`
		Name     string `yaml:"name"`
		Engine   string `yaml:"engine"`
	}

	Application struct {
		Port        string `yaml:"port"`
		ClusterName string `yaml:"cluster_name"`
	}
}

func NewConfig() (*Config, error) {
	cfgFile, err := ioutil.ReadFile("./data/config.yml")

	if err != nil {
		application_loger.LogError(err.Error())
		return nil, err
	}

	conf := Config{}

	err = yaml.Unmarshal(cfgFile, &conf)

	if err != nil {
		application_loger.LogError(err.Error())
		return nil, err
	}

	return &conf, nil
}
