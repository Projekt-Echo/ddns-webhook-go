package routes

import (
	"ddns-webhook-go/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterWebhookRoute(app *fiber.App) {
	app.Post("/webhook", controller.WebhookController)
}
