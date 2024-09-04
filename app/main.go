package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gofiber-template/internal/config"
	"gofiber-template/internal/database"
	"gofiber-template/internal/logger"
)

func init() {
	logger.InitLogger()
	config.LoadConfig()
	database.InitDatabase()
}
func main() {
	fiberConfig := fiber.Config{
		AppName:     viper.GetString("app.name"),
		Prefork:     viper.GetBool("app.prefork"),
		JSONDecoder: sonic.Unmarshal,
		JSONEncoder: sonic.Marshal,
	}

	app := fiber.New(fiberConfig)
	if err := app.Listen(viper.GetString("app.addr")); err != nil {
		logger.Logger.Sugar().Fatalf("app failed to start: %v", err)
	}
}
