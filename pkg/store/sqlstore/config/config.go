package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

func (db *DBConfig) GetConnactionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}

func GetDataBaseConfig(fileName string, path string) (*DBConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg DBConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
