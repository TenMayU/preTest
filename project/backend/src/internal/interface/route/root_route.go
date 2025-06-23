package routes

import (
	middleware "backendrest/src/http"
	"backendrest/src/internal/domain/quote"
	"backendrest/src/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app *fiber.App, authService user.AuthService, quoteSerivce quote.QuoteService) {
	api := app.Group("/api")
	RegisterAuthRoutes(api.Group("/auth"), authService)
	app.Use(middleware.JwtMiddleware)
	RegisterQuoteRoutes(api.Group("/quote"), quoteSerivce)
}
