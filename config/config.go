package config

import (
	"os"
	"strconv"
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
		Host       string `toml:"host"`
		Port       int    `toml:"port"`
		Username   string `toml:"username"`
		Password   string `toml:"password"`
		Database   string `toml:"database"`
		Connection string `toml:"connection"`
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

	defaultConfig.App.Port, _ = strconv.Atoi(os.Getenv("PORT"))

	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Connection = os.Getenv("DATABASE_URL")
	defaultConfig.Database.Host = os.Getenv("DB_HOST")
	defaultConfig.Database.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	defaultConfig.Database.Database = os.Getenv("DB_DATABASE")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

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
