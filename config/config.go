package config

import (
	"sync"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Env    string `toml:"env"`
		Port   int    `toml:"port"`
		JWTKey string `toml:"jwtkey"`
	} `toml:"app"`
	Database struct {
		Driver string `toml:"driver"`
		// DBURL  string `toml:"dburl"`
		Address  string `toml:"address"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.App.Port = 5006

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("../config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("error when load config file", err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)

	if err != nil {
		log.Info("error when parse config file", err)
		return &defaultConfig
	}
	return &finalConfig
}
