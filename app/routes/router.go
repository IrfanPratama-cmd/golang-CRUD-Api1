package routes

import (
	"final-project/app/controller"
	"final-project/app/model"
	"final-project/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Handler(app *fiber.App) {
	app.Use(cors.New())

	services.InitDatabase()
	db := services.DB
	db.AutoMigrate(&model.Car{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world"})
	})

	app.Get("/api/v1/cars", controller.GetCar)
	app.Post("/api/v1/cars", controller.PostCar)
	app.Put("/api/v1/cars/:id", controller.PutCar)
	app.Get("/api/v1/cars/:id", controller.GetCarID)
	app.Delete("/api/v1/cars/:id", controller.DeleteCar)
}
