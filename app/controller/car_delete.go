package controller

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//DeleteCar godoc
// @Summary Delete Car
// @Description Delete Car
// @Tags Car
// @Accept application/json
// @Produce application/json
// @Security TokenKey
// @Param id   path      int  true  "Car ID"
// @Success 200 {string} success "OK"
// @Router /cars/{id} [delete]

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	var car model.Car
	result := db.Where(`id = ?`, id).First(&car)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	result.Delete(&car)

	message := fmt.Sprintf(`car with id %s has been deleted`, id)

	return c.JSON(fiber.Map{
		"message": message,
	})
}
