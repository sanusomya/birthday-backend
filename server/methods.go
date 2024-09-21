package server

import (
	"birthday/birthday"
	"birthday/database"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func registerEndpoints(e *echo.Echo) {
	editGroup := e.Group("/edit")
	editGroup.PUT("/", updateBirthday, validationMiddleware)
	editGroup.PUT("/name", updateBirthdayName, nameValidationMiddleware)
	editGroup.PUT("/number", updateBirthdayNumber, mobileValidationMiddleware)
	editGroup.PUT("/date", updateBirthdayDate, dateValidationMiddleware)
	e.GET("/", getAllBirthdays)
	e.GET("/today", getAllBirthdaysForToday)
	e.GET("/month", getAllBirthdaysForThisMonth)
	e.DELETE("/", deleteBirthday, validationMiddleware)
	e.POST("/", addBirthday, validationMiddleware)
}

var _ = godotenv.Load()
var uri = os.Getenv("db_url")
var db = os.Getenv("database")
var collection = os.Getenv("db_coll")
var coll, _ = database.ConnectDB(uri, db, collection)

func getAllBirthdays(c echo.Context) error {
	birthdays, err := database.GetAll(coll)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, birthdays)
}

func addBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	var temp = birthday.Birthday{}
	temp = body
	err = database.Add(coll, temp)
	if err != nil {
		c.JSON(http.StatusFound, err.Error())
		return err
	}
	return c.JSON(http.StatusCreated, temp)
}

func deleteBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	var temp = birthday.Birthday{}
	temp = body
	err = database.Delete(coll, temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, temp)
}

func updateBirthday(c echo.Context) error {
	var body birthday.Birthday
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	params := c.QueryParams()

	var temp = birthday.Birthday{}
	temp = body
	mobile, err := strconv.Atoi(params["mobile"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	err = database.Edit(coll, params["name"][0], int64(mobile), temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	return c.JSON(http.StatusPartialContent, temp)
}

func updateBirthdayName(c echo.Context) error {
	var body string
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	params := c.QueryParams()
	var temp = birthday.Birthday{}
	mobile, err := strconv.Atoi(params["mobile"][0])
	name := params["name"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, "here")
		return err
	}
	bday, err := database.FindByNameAndMobile(coll, int64(mobile), name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	temp.Name = body
	temp.Date = bday.Date
	temp.Month = bday.Month
	temp.Mobile = int64(mobile)
	err = database.Edit(coll, name, int64(mobile), temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	return c.JSON(http.StatusPartialContent, temp)
}

func updateBirthdayNumber(c echo.Context) error {
	var body int64
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	params := c.QueryParams()
	var temp = birthday.Birthday{}
	mobile, err := strconv.Atoi(params["mobile"][0])
	name := params["name"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, "here")
		return err
	}
	bday, err := database.FindByNameAndMobile(coll, int64(mobile), name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	temp.Name = name
	temp.Date = bday.Date
	temp.Month = bday.Month
	temp.Mobile = body
	err = database.Edit(coll, name, int64(mobile), temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	return c.JSON(http.StatusPartialContent, temp)
}

func updateBirthdayDate(c echo.Context) error {
	var body struct {
		Date  int
		Month string
	}
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind body")
		return err
	}
	params := c.QueryParams()
	var temp = birthday.Birthday{}
	mobile, err := strconv.Atoi(params["mobile"][0])
	name := params["name"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	temp.Name = name
	temp.Date = int8(body.Date)
	temp.Month = body.Month
	temp.Mobile = int64(mobile)
	err = database.Edit(coll, name, int64(mobile), temp)
	if err != nil {
		fmt.Println("here")
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	return c.JSON(http.StatusPartialContent, temp)
}

func getAllBirthdaysForToday(c echo.Context) error {
	params := c.QueryParams()
	date, err := strconv.Atoi(params["date"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	birthdays, err := database.FindForToday(coll, params["month"][0], int8(date))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, birthdays)
}

func getAllBirthdaysForThisMonth(c echo.Context) error {
	params := c.QueryParams()
	birthdays, err := database.FindForThisMonth(coll, params["month"][0])
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, birthdays)
}
