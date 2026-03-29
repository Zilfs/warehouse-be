package app

import "github.com/gofiber/fiber/v3"

func RunServer() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server start")
	})
	app.Listen(":8000")
}
