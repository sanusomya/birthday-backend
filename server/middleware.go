package server

import (
	"birthday/birthday"
	"birthday/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func checkParams(bday birthday.Birthday) bool {
	validDate := utils.CheckDates(bday.Date, bday.Month)
	validName := utils.ValidName(bday.Name)
	validMobile := utils.ValidMobile(bday.Mobile)
	return validDate && validMobile && validName
}

func validationMiddleware(m echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bday := birthday.Birthday{}
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot read request body")
			return err
		}
		err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&bday)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot decode request body")
			return err
		}
		isValid := checkParams(bday)
		if !isValid {
			c.JSON(http.StatusBadRequest, map[string]string{
				"name":   "should only contains alphabets with length less than 10",
				"mobile": "should contain only numbers with length 10",
				"date":   "month date combination be according to the calender, month shoud be 3 letter initial of the full name.",
			})
			return nil
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return m(c)
	}
}

func nameValidationMiddleware(m echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var name string
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot read request body")
			return err
		}
		err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&name)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}
		isValid := utils.ValidName(name)
		if !isValid {
			c.JSON(http.StatusBadRequest, map[string]string{
				"name": "should only contains alphabets with length less than 10",
			})
			return nil
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return m(c)
	}
}

func mobileValidationMiddleware(m echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var mobile int64
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot read request body")
			return err
		}
		err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&mobile)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}
		isValid := utils.ValidMobile(mobile)
		if !isValid {
			c.JSON(http.StatusBadRequest, map[string]string{
				"mobile": "should contain only numbers with length 10",
			})
			return nil
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return m(c)
	}
}

func dateValidationMiddleware(m echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var bday struct {
			Date  int8
			Month string
		}
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot read request body")
			return err
		}
		err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&bday)
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot decode request body")
			return err
		}
		isValid := utils.CheckDates(bday.Date, bday.Month)
		if !isValid {
			c.JSON(http.StatusBadRequest, map[string]string{
				"date": "month date combination be according to the calender, month shoud be 3 letter initial of the full name.",
			})
			return nil
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return m(c)
	}
}
