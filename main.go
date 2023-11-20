package main

import (
	sql "database/sql"
	fmt "fmt"

	echo "github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var cars []Car
var db *sql.DB

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

func execQuery(query string, params ...any) error {
	database, err := getDb()
	if err != nil {
		return err
	}
	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		return err
	}
	return nil
}

func getDb() (*sql.DB, error) {
	if db != nil {
		return db, nil
	} else {
		database, err := sql.Open("sqlite3", "cars.db")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer db.Close()

		db = database
		return database, nil
	}
}

func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Start(":8080")
}

func saveCar(car Car) error {
	err := execQuery("INSERT INTO cars (name, price)  values ($1, $2)", car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
