package main

import (
	sql "database/sql"

	echo "github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var cars []Car

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
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

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cars (name,price) VALUES ($1,$2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
