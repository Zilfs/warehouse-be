package app

import (
	"log"
	"warehouse/config"

	"github.com/gofiber/fiber/v3"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnetionPostgres()
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
		return
	}

	_ = db

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server start")
	})
	app.Listen(":8000")
}
