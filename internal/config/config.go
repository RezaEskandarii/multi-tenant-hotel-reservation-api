package config

import (
	"gopkg.in/yaml.v2"
	"hotel-reservation/pkg/applogger"
	"io/ioutil"
	"os"
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
		Port            string `yaml:"port"`
		ClusterName     string `yaml:"cluster_name"`
		IgnoreMigration bool   `yaml:"ignore_migration"`
		SqlDebug        bool   `json:"sql_debug"`
	}
}

func NewConfig() (*Config, error) {
	path := "./resources/config.yml"

	logger := applogger.New()

	if os.Getenv("CONFIG_PATH") != "" {

		path = os.Getenv("CONFIG_PATH")
	}

	cfgFile, err := ioutil.ReadFile(path)

	if err != nil {
		logger.LogError(err.Error())
		return nil, err
	}

	conf := Config{}

	err = yaml.Unmarshal(cfgFile, &conf)

	if err != nil {
		logger.LogError(err.Error())
		return nil, err
	}

	return &conf, nil
}
