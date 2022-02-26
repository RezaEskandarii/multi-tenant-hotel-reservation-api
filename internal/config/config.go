package config

import (
	"gopkg.in/yaml.v2"
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

	// Application is application config
	Application struct {
		Port            string `yaml:"port"`
		ClusterName     string `yaml:"cluster_name"`
		IgnoreMigration bool   `yaml:"ignore_migration"`
		SqlDebug        bool   `yaml:"sql_debug"`
	}
	// Minio file management
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		UseSSL          bool   `yaml:"use_ssl"`
	}
	MessageBroker struct {
		Url string `yaml:"url"`
	}
}

// NewConfig reads config from yml file copies to config struct and returns config struct
func NewConfig() (*Config, error) {

	path := "./resources/config.yml"

	if os.Getenv("CONFIG_PATH") != "" {
		path = os.Getenv("CONFIG_PATH")
	}

	cfgFile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	conf := Config{}

	err = yaml.Unmarshal(cfgFile, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
