package config

import (
	"github.com/spf13/viper"
	"gofiber-template/internal/logger"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Logger.Sugar().Fatal("%s", err)
	}
}
