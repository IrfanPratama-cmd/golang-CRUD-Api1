package controller

import (
	"errors"
	"final-project/app/model"
	"final-project/app/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PutCar godoc
// @Summary Update Car
// @Description Update Car
// @Tags Car
// @Accept application/json
// @Produce application/json
// @Security TokenKey
// @Param id   path      int  true  "Car ID"
// @Param car body model.CarAPI true "Car body"
// @Success 200 {object} model.Car{} "OK"
// @Success 400
// @Success 404
// @Router /cars/{id} [put]

func PutCar(c *fiber.Ctx) error {
	var carApi model.CarApi

	id := c.Params("id")

	var car model.Car

	db := services.DB

	if err := db.Model(&model.Car{}).First(&car, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := c.BodyParser(&carApi); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	update := &model.Car{CarApi: carApi}
	db.Model(&model.Car{}).Where(`id = ?`, id).Updates(update)
	message := fmt.Sprintf(`car with id %s has been updated`, id)

	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    update,
	})
}
