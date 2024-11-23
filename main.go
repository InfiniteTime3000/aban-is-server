package main

import (
	"sync"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	Msg := "..."
	var mu sync.Mutex

	app := fiber.New()
	app.Use(cors.new(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "Origin, Content-Type, Accept"
	}))

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
			"msg": Msg,
		})
	})

	app.Listen(":6969")
}
