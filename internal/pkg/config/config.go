package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(fileConfigPath string) (Config, error) {
	config := Config{}

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
