package bithday_handler

import "github.com/labstack/echo/v4"

func RegisterEndpoints(e *echo.Echo) {
	editGroup := e.Group("/edit")
	editGroup.PUT("/", updateBirthday, validationMiddleware)
	editGroup.PUT("/name", updateBirthdayName, nameValidationMiddleware)
	editGroup.PUT("/number", updateBirthdayNumber, mobileValidationMiddleware)
	e.GET("/", getAllBirthdays)
	e.GET("/today", getAllBirthdaysForToday)
	e.GET("/month", getAllBirthdaysForThisMonth)
	e.DELETE("/", deleteBirthday, validationMiddleware)
	e.POST("/", addBirthday, validationMiddleware)
}
