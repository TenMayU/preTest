package main

import (
	authService "backendrest/src/internal/appication/auth"
	quoteService "backendrest/src/internal/appication/quote"
	"backendrest/src/internal/domain/quote"
	"backendrest/src/internal/domain/user"
	infrastructure "backendrest/src/internal/infrastructure/db"
	repository "backendrest/src/internal/infrastructure/repository"
	routes "backendrest/src/internal/interface/route"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")
	log.Println(port)
	db := infrastructure.NewSQLiteDB()
	err := db.AutoMigrate(&user.User{}, &quote.Quote{}, &quote.QuoteVoting{})
	for _, e := range os.Environ() {
		log.Println(e)
	}
	if err != nil {
		fmt.Println("Migration failed:", err)
	}
	authRepo := repository.NewAuthService(db)
	authService := authService.NewAuthService(authRepo)

	quoteRepo := repository.NewQuoteService(db)
	quoteSerivce := quoteService.NewAuthService(quoteRepo)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // หรือ "*" ก็ได้
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.RegisterAllRoutes(app, authService, quoteSerivce)
	app.Listen(":" + port)
}
