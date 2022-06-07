package main

import (
	"final-project/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @title Golang Mini Bootcamp 2022 - By Irfan Pratama
// @version 1.0.0
// @description API Documentation
// @host localhost:8081
// @schemes http
// @BasePath /api/v1

// @in header

func main() {
	app := fiber.New()

	routes.Handler(app)
	log.Fatal(app.Listen(":8081"))
}
