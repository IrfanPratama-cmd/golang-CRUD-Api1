package controller

import (
	"final-project/app/model"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetCar(t *testing.T) {

	app := fiber.New()
	app.Get("/cars", GetCar)

	db := services.InitDatabaseTest()
	body := model.Car{
		CarApi: model.CarApi{
			CarName:     "Veyron",
			Brand:       "Bugati",
			Year:        2022,
			Price:       100000000,
			Description: "Mobil Apik",
		},
	}
	db.Create(&body)

	req, _ := http.NewRequest("GET", "/movies", nil)
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get request")
}
