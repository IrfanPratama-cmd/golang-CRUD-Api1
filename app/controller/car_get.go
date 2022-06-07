package controller

import (
	"final-project/app/model"
	"final-project/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetCar godoc
// @Summary Get list of Car
// @Description Get list of Car
// @Accept application/json
// @Produce application/json
// @Security TokenKey
// @Success 200 {object} []model.Car{} "Ok"
// @Router /cars [get]
// @Tags Car

func GetCar(c *fiber.Ctx) error {
	var car []model.Car
	db := services.DB
	db.Model(&model.Car{}).Find(&car)
	return c.Status(200).JSON(car)
}
