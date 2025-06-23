package routes

import (
	"backendrest/src/internal/domain/user"
	"backendrest/src/internal/interface/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(r fiber.Router, service user.AuthService) {
	handler := handler.NewHttpAuth(service)
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)
	r.Post("/test", handler.Test)
}
