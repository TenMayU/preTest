package routes

import (
	"backendrest/src/internal/domain/quote"
	"backendrest/src/internal/interface/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterQuoteRoutes(r fiber.Router, service quote.QuoteService) {
	handler := handler.NewHttpQuote(service)
	r.Get("/getAllQuote", handler.GetAllQuote)
	r.Get("/getAllVoted", handler.GetAllVoted)
	r.Post("/create", handler.Create)
	r.Post("/voting", handler.Voting)
	r.Post("/search", handler.GetQuoteBySearch)
	r.Post("/update", handler.Update)
}
