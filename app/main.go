package main

import (
	"ddns-webhook-go/internal/config"
	"ddns-webhook-go/internal/database"
	"ddns-webhook-go/internal/logger"
	"ddns-webhook-go/internal/routes"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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

	routes.RegisterWebhookRoute(app)

	if err := app.Listen(viper.GetString("app.addr")); err != nil {
		logger.Logger.Sugar().Fatalf("app failed to start: %v", err)
	}

}
