package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB       DB
	ApiKey   string `mapstructure:"api_key"`
	LogLevel string `mapstructure:"log_level"`
}

type DB struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
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

	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := cfg.ParseEnv(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
func (c *Config) ParseEnv() error {

	return nil
}
func (db *DB) GetConnactionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}
