package server

import (
	"birthday/birthday"
	"birthday/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var coll, err = database.ConnectDB()

func registerEndpoints(e *echo.Echo) {
	e.GET("/", getAllBirthdays)
	e.POST("/", addBirthday)
	e.DELETE("/", deleteBirthday)
	e.PUT("/", updateBirthday)
}
func getAllBirthdays(c echo.Context) error {
	if err != nil {
		panic(err)
	}
	birthdays, err := database.GetAll(coll)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, birthdays)
}

func addBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}
	var temp = birthday.Birthday{}
	temp = body
	err = database.Add(coll, temp)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, temp)
}

func deleteBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}
	var temp = birthday.Birthday{}
	temp = body
	err = database.Delete(coll, temp)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, temp)
}

func updateBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}
	params := c.QueryParams()

	var temp = birthday.Birthday{}
	temp = body
	mobile, err := strconv.Atoi(params["mobile"][0])
	if err != nil {
		panic(err)
	}
	err = database.Edit(coll, params["name"][0], int64(mobile), temp)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusPartialContent, temp)
}
