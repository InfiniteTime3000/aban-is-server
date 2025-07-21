package main

import (
	"fmt"
	"sync"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const msgFile = "msg.txt"

func loadMessage() string {
	data, err := os.ReadFile(msgFile)

	if err != nil {
		fmt.Println("No existing message file or failed to read, using default message.")
		return "sleeping or his laptop is off (server is down)"
	}
	return string(data)
}

func saveMessage(msg string) {
	err := os.WriteFile(msgFile, []byte(msg), 0644)
	fmt.Println("Saving message:", msg)

	if err != nil {
		fmt.Println("Failed to save message:", err)
	}
}

func main() {
	Msg := loadMessage()
	var mu sync.Mutex

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "Origin, Content-Type, Accept",
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
		saveMessage(Msg)

		mu.Unlock()
		return c.JSON(fiber.Map{
			"msg": Msg,
		})
	})

	app.Listen(":6969")
}