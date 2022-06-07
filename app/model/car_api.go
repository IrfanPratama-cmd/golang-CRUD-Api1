package model

type CarApi struct {
	CarName     string `json:"car-name,omitempty" example:"Mclaren"`
	Brand       string `json:"brand,omitempty" example:"Mercedes"`
	Year        int    `json:"year,omitempty" example:"2020"`
	Price       int    `json:"price,omitempty" example:"10000"`
	Description string `json:"description,omitempty" example:"Good"`
}
