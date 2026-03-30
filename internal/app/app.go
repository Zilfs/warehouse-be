package app

import (
	"fmt"
	"log"
	"warehouse/config"

	"github.com/gofiber/fiber/v3"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	_ = db

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server start")
	})

	port := cfg.App.AppPort
	if port == "" {
		port = "8000"
	}

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
