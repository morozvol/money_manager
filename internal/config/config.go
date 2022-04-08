package config

import (
	"github.com/morozvol/money_manager/pkg/store/sqlstore/config"
	"github.com/spf13/viper"
)

type Config struct {
	DB       *config.DBConfig
	ApiKey   string `mapstructure:"api_key"`
	LogLevel string `mapstructure:"log_level"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	dateBaseConfig, err := config.GetDataBaseConfig("db")
	if err != nil {
		return nil, err
	}
	cfg.DB = dateBaseConfig

	return &cfg, nil
}
