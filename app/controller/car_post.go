package controller

import (
	"final-project/app/model"
	"final-project/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostCar godoc
// @Summary Create new Car
// @Description Create new Car
// @Accept application/json
// @Produce application/json
// @Security TokenKey
// @Success 200 {object} []model.Car{} "Created"
// @Success 400
// @Param data body model.CarApi true "Car data"
// @Router /cars [post]
// @Tags Car

func PostCar(c *fiber.Ctx) error {
	var carApi model.CarApi

	db := services.DB

	if err := c.BodyParser(&carApi); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	car := &model.Car{CarApi: carApi}
	db.Model(&model.Car{}).Create(car)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    car,
	})
}
