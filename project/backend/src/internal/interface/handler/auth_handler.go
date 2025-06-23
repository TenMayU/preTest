package handler

import (
	"backendrest/src/internal/domain/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	handler user.AuthService
}

func NewHttpAuth(service user.AuthService) *AuthHandler {
	return &AuthHandler{handler: service}
}

func (s *AuthHandler) Login(c *fiber.Ctx) error {
	var bodyParse user.User
	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"status":  false,
		})
	}

	user, ok, err := s.handler.Login(c.Context(), bodyParse.UserName, bodyParse.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  false,
		})
	}

	if ok == "" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user login successfully",
		"user":    user,
		"status":  true,
		"token":   ok,
	})

}

func (s *AuthHandler) Register(c *fiber.Ctx) error {

	var bodyParse user.User

	if err := c.BodyParser(&bodyParse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	ok, err := s.handler.Register(c.Context(), bodyParse.UserName, bodyParse.Password, bodyParse.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register has fail",
			"status":  false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
		"status":  true,
	})

}

func (s *AuthHandler) Test(c *fiber.Ctx) error {
	type request struct {
		Text string `json:"Text"`
	}
	// อ่าน raw body จาก context
	rawBody := c.Body() // []byte

	// แปลงเป็น string เพื่อ log หรือดูข้อความ
	payloadText := string(rawBody)

	// ตัวอย่าง: แค่ log หรือส่งกลับไปดู
	fmt.Println("Raw payload:", payloadText)
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid JSON")
	}

	return c.SendString("Received: " + body.Text)

}
