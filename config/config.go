package config

import (
	"bytes"
	"fmt"
	"runtime"

	_ "embed"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		ENV      string `mapstructure:"env"`
		APPNAME  string `mapstructure:"app_name"`
		LOGLEVEL string `mapstructure:"log_level"`
		SECRET   string `mapstructure:"secret"`
	} `mapstructure:"app"`
	Server struct {
		HTTPAddress string `mapstructure:"http_address"`
	} `mapstructure:"server"`
	Database struct {
		Host        string `mapstructure:"host"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		Database    string `mapstructure:"database"`
		MaxIdleConn int    `mapstructure:"max_idle"`
		MaxOpenConn int    `mapstructure:"max_open"`
	} `mapstructure:"database"`
}

//go:embed default.yaml
var defaultData []byte

var Cfg *Config

func init() {
	Cfg = LoadConfig()
}

func LoadConfig() *Config {
	// Initialize viper
	cfg := &Config{}

	v := viper.New()

	_, filePath, _, _ := runtime.Caller(0)
	configFile := filePath[:len(filePath)-9]

	v.SetConfigFile(configFile + "config" + ".yaml")

	err := v.ReadConfig(bytes.NewBuffer(defaultData))
	if err != nil {
		panic(err)
	}

	// Load the config file
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// Unmarshal the config into a struct
	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))
	}

	// Use the config values
	return cfg
}
