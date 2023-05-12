package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver               string        `mapstructure:"DB_DRIVER"`
	DBConnection           string        `mapstructure:"DB_CONNECTION"`
	ServerPort             string        `mapstructure:"SERVER_PORT"`
	JwtAccessTokenKey      string        `mapstructure:"JWT_ACCESS_TOKEN_KEY"`
	JwtAccessTokenDuration time.Duration `mapstructure:"JWT_ACCESS_TOKEN_DURATION"`
	JwtRefreshTokenKey     string        `mapstructure:"JWT_REFRESH_TOKEN_KEY"`
	RefreshTokenDuration   time.Duration `mapstructure:"JWT_REFRESH_TOKEN_DURATION"`
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
