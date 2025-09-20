package bithday_handler

import (
	"strconv"

	birthday "github.com/sanusomya/birthday-backend/models"

	"net/http"
	"os"

	"github.com/sanusomya/birthday-backend/database"

	// "strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var _ = godotenv.Load()
var table = os.Getenv("table")
var coll = database.ConnectDB()

func getAllBirthdays(c echo.Context) error {
	birthdays, err := database.GetAll(coll, table)
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
	err = database.Add(coll, table, temp)
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
	err = database.Delete(coll, table, temp)
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
	err = database.Edit(coll, table, params["name"][0], int64(mobile), temp)
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
	mobile, err := strconv.Atoi(params["mobile"][0])
	name := params["name"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot convert to integer")
		return err
	}
	bday, err := database.Get(coll, table, name, int64(mobile))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	err = database.Delete(coll, table, bday)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	temp := birthday.Birthday{
		Person:     body,
		Cell:       bday.Cell,
		Birthdate:  bday.Birthdate,
		Birthmonth: bday.Birthmonth,
	}
	err = database.Add(coll, table, temp)
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
	mobile, err := strconv.Atoi(params["mobile"][0])
	name := params["name"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot convert to integer")
		return err
	}
	bday, err := database.Get(coll, table, name, int64(mobile))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	err = database.Delete(coll, table, bday)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	temp := birthday.Birthday{
		Person:     bday.Person,
		Cell:       body,
		Birthdate:  bday.Birthdate,
		Birthmonth: bday.Birthmonth,
	}
	err = database.Add(coll, table, temp)
	if err != nil {
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
	birthdays, err := database.FindForToday(coll, table, params["month"][0], int8(date))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, birthdays)
}

func getAllBirthdaysForThisMonth(c echo.Context) error {
	params := c.QueryParams()
	birthdays, err := database.FindForThisMonth(coll, table, params["month"][0])
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, birthdays)
}
