package main

import (
	"sync"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	Msg := "..."
	var mu sync.Mutex

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		mu.Lock()
		defer mu.Unlock()
		return c.JSON(fiber.Map{
			"msg": Msg,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		type Request struct {
			newMsg string `json:"msg"`
		}
		var body Request
		fmt.Println(string(c.Body()))
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}
		mu.Lock()
		Msg = body.newMsg
		mu.Unlock()
		return c.JSON(fiber.Map{
			"msg": body.newMsg,
		})
	})

	app.Listen(":6969")
}
