package config

import (
	"ddns-webhook-go/internal/logger"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Logger.Sugar().Fatal("%s", err)
	}
}
