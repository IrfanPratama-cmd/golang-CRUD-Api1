package controller

import (
	"bytes"
	"final-project/app/model"
	"final-project/app/services"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutCar(t *testing.T) {
	app := fiber.New()
	app.Put("/cars/:id", PutCar)

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
	id := strconv.Itoa(*&body.ID)

	sampleNotFound := model.Car{
		CarApi: model.CarApi{
			CarName:     "Veyron",
			Brand:       "Bugati",
			Year:        2022,
			Price:       100000000,
			Description: "Mobil Apik",
		},
	}

	db.Create(&sampleNotFound)
	idNotFound := strconv.Itoa(*&sampleNotFound.ID)
	db.Where(`id = ?`, idNotFound).Delete(&sampleNotFound)

	// positive case
	req, _ := http.NewRequest("PUT", "/cars/"+id, payload)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code success")

	// case error param id
	req, _ = http.NewRequest("PUT", "/cars/string", payload)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad request")

	// case error body parser
	req, _ = http.NewRequest("PUT", "/cars/"+id, nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad request")

	// case error id not exist
	req, _ = http.NewRequest("PUT", "/cars/"+idNotFound, payload)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "response id not exist")
}
