package controller

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type payload struct {
	Key      string `json:"key"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Group    string `json:"group"`
	Priority int    `json:"priority"`
}

func WebhookController(c *fiber.Ctx) error {
	req := payload{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	form := url.Values{}
	form.Set("key", req.Key)
	form.Set("content", req.Content)
	form.Set("type", req.Type)
	form.Set("title", req.Title)
	form.Set("group", req.Group)
	form.Set("priority", strconv.Itoa(req.Priority))

	resp, err := http.Post("https://push.loliconapp.top/v2/api/message/push", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot read response body",
		})
	}

	return c.Status(fiber.StatusOK).SendString(string(body))
}
