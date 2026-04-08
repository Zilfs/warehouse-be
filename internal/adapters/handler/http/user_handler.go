package http

import (
	"warehouse/internal/core/domain/entity"
	"warehouse/internal/core/domain/model"
	"warehouse/internal/core/ports"
	"warehouse/pkg/utils"

	"github.com/gofiber/fiber/v3" // Pastikan v3 konsisten
)

type UserHandler struct {
	usecase ports.UserUsecase
}

func NewUserHandler(app *fiber.App, uc ports.UserUsecase) {
	h := &UserHandler{usecase: uc}

	api := app.Group("/api/v1")

	// SALAH: api.Post("/users", uc.CreateUser)
	// BENAR: Gunakan method dari struct handler (h.Create)
	api.Post("/users", h.Create)
	api.Get("/users", h.GetAll)
	api.Get("/users/:id", h.GetByID)
	api.Put("/users/:id", h.Update)

}

func (h *UserHandler) Create(c fiber.Ctx) error {
	var req model.UserRequest

	// 1. Parse JSON body ke struct request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// 2. Mapping dari Model Request ke Entity
	userEntity := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	// 3. Panggil Usecase (Logic)
	if err := h.usecase.CreateUser(c.Context(), userEntity); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	users, err := h.usecase.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": users})
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	userId, err := utils.StringToInt(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.usecase.GetUserByID(c.Context(), userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": user})
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	var req model.UserRequest
	id := c.Params("id")
	userId, err := utils.StringToInt(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.usecase.GetUserByID(c.Context(), userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password

	if err := h.usecase.UpdateUser(c.Context(), user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User updated successfully"})
}
