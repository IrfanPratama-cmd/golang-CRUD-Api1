package controller

import (
	"bytes"
	"final-project/app/model"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostCar(t *testing.T) {
	app := fiber.New()
	app.Post("/cars", PostCar)

	payload := bytes.NewReader([]byte(`
	{ 
		"car-name" : "Ferrari",
		"brand" : "Ferrari",
		"year" : 2010,
		"price" : 2000000,
		"description" : "Apik"
  
	}
	`))

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

	// positive case
	req, _ := http.NewRequest("POST", "/movies", payload)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code success")

	// case error body parser
	req, _ = http.NewRequest("POST", "/movies", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad request")
}
