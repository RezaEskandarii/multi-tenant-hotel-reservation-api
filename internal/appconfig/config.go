package appconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config application Config struct
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

	// Application is application Config
	Application struct {
		Port               string `yaml:"port"`
		ClusterName        string `yaml:"cluster_name"`
		IgnoreMigration    bool   `yaml:"ignore_migration"`
		DebugMode          bool   `yaml:"debug_mode"`
		MetricEndPointPort int    `yaml:"metric_end_point_port"`
		AllowedOrigins     string `yaml:"allowed_origins"`
	}
	// Minio file management
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		UseSSL          bool   `yaml:"use_ssl"`
	}
	// Message broker Config
	MessageBroker struct {
		Url string `yaml:"url"`
	}
	// Smtp Config
	Smtp struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	}
	Authentication struct {
		JwtKey         string `yaml:"jwt_key"`
		TokenAliveTime int    `yaml:"token_alive_time"` // minute
	}

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		CacheDB  int    `yaml:"cache_db"`
	}
}

// New reads Config from yml file copies to Config struct and returns Config struct
func New() (*Config, error) {

	path := "./resources/Config.yml"

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
