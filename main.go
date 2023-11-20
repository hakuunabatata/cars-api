package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}

	cars = append(cars, *car)

	return c.JSON(200, cars)

}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Start(":8080")
}
