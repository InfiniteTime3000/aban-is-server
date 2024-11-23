package main

import (
	"sync"
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
		mu.Lock()
		Msg = string(c.Body())
		mu.Unlock()
		return c.JSON(fiber.Map{
			"msg": body.newMsg,
		})
	})

	app.Listen(":6969")
}
