package handler

import (
	"backendrest/src/internal/domain/quote"

	"github.com/gofiber/fiber/v2"
)

type QuoteHandler struct {
	handler quote.QuoteService
}

func NewHttpQuote(service quote.QuoteService) *QuoteHandler {
	return &QuoteHandler{handler: service}
}

func (s *QuoteHandler) GetAllVoted(c *fiber.Ctx) error {
	/* var quote []quote.QuoteVoting */
	ok, err := s.handler.GetAllVoted(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "get successfully",
		"data":    ok,
	})
}

func (s *QuoteHandler) GetAllQuote(c *fiber.Ctx) error {
	/* var quote []quote.Quote */
	ok, err := s.handler.GetAllQuote(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "get successfully",
		"data":    ok,
	})
}
func (s *QuoteHandler) GetQuoteBySearch(c *fiber.Ctx) error {

	var bodyParse quote.Quote
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"status":  false,
		})
	}

	ok, err := s.handler.GetQuoteBySearch(c.Context(), bodyParse.Text)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "search successfully",
		"status":  true,
		"data":    ok,
	})
}
func (s *QuoteHandler) Create(c *fiber.Ctx) error {
	var bodyParse quote.Quote
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"status":  false,
		})
	}

	if bodyParse.Text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "all fields are required",
			"status":  false,
		})
	}

	ok, err := s.handler.Create(c.Context(), bodyParse.Text)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "create successfully",
		"status":  true,
	})
}
func (s *QuoteHandler) Voting(c *fiber.Ctx) error {
	var bodyParse quote.QuoteVoting
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"status":  false,
		})
	}

	if bodyParse.QuoteId == 0 || bodyParse.User == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "all fields are required",
			"status":  false,
		})
	}

	ok, mgs, err := s.handler.Voting(c.Context(), bodyParse.QuoteId, bodyParse.User)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": mgs,
			"status":  false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "voted successfully",
		"status":  true,
	})
}

func (s *QuoteHandler) Update(c *fiber.Ctx) error {
	var bodyParse quote.Quote
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid request body",
			"status": false,
		})
	}

	ok, err := s.handler.Update(c.Context(), bodyParse.ID, bodyParse.Text)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "update has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "update successfully",
		"status":  true,
	})
}
