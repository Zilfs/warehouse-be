package app

import (
	"fmt"
	"log"
	"warehouse/config"
	"warehouse/internal/adapters/handler/http"        // Import handler Anda
	"warehouse/internal/adapters/repository/postgres" // Import repo Anda
	"warehouse/internal/core/usecase"

	"github.com/gofiber/fiber/v3"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	// 1. Inisialisasi Fiber
	app := fiber.New()

	// 2. Inisialisasi Layer (Dependency Injection)
	// Pastikan konstruktor New ini sesuai dengan yang Anda tulis di masing-masing file
	userRepo := postgres.NewUserRepository(db.DB)
	userUC := usecase.NewUserUsecase(userRepo)

	// 3. Daftarkan Handler ke Fiber (Ini yang memperbaiki 404)
	http.NewUserHandler(app, userUC)

	// Route testing bawaan Anda
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server start")
	})

	port := cfg.App.AppPort
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server is running on port", port)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
