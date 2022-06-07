package controller

import (
	"final-project/app/model"
	"final-project/app/services"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetCarID(t *testing.T) {

	app := fiber.New()
	app.Get("/cars/:id", GetCarID)

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
	id := strconv.Itoa(*&body.ID)

	// positive case
	req, _ := http.NewRequest("GET", "/cars/"+id, nil)
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code success")

	// case error param id
	req, _ = http.NewRequest("GET", "/cars/string", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad request")

	// case error id not exist
	req, _ = http.NewRequest("GET", "/cars/0", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "response id not exist")

}
