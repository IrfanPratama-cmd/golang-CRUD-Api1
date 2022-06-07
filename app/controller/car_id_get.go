package controller

import (
	"final-project/app/model"
	"final-project/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetMovieID get cars by id
// @Summary get cars by id
// @Description get cars by id
// @Tags Car
// @Accept application/json
// @Produce application/json
// @Security TokenKey
// @Param id   path      int  true  "Car ID"
// @Success 200 {object} model.Car{} "OK"
// @Success 400
// @Success 404
// @Router /cars/{id} [get]

func GetCarID(c *fiber.Ctx) error {
	id := c.Params("id")

	db := services.DB

	var car model.Car
	result := db.Where(`id = ?`, id).First(&car)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	return c.JSON(car)
}
