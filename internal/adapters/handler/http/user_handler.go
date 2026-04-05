package http

import (
	"warehouse/internal/core/domain/entity"
	"warehouse/internal/core/domain/model"
	"warehouse/internal/core/ports"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	usecase ports.UserUsecase
}

func NewUserHandler(app *fiber.App, uc ports.UserUsecase) {
	h := &UserHandler{usecase: uc}
	api := app.Group("/api/v1")

	api.Post("/users", h.usecase.CreateUser)
	api.Get("/users", h.usecase.GetAllUsers)
	api.Post("/users:ID", h.usecase.GetUserByID)
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	var req model.UserRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.usecase.CreateUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	users, _ := h.usecase.GetAllUsers(c.Context())
	var response []model.UserResponse
	for _, u := range users {
		response = append(response, model.UserResponse{ID: u.ID, Username: u.Username, Email: u.Email})
	}
	return c.JSON(response)
}
