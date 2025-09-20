package server

import (
	"os"
	"github.com/labstack/echo/v4"
	"github.com/sanusomya/birthday-backend/handlers/birthdays"
)

func MyEchoServer() {
	var e = echo.New()
	bithday_handler.RegisterEndpoints(e)

	port := os.Getenv("birthday_app_port")
	if port == ""{
		port = "8002"
	}
	e.Logger.Error(e.Start(":"+port))
}
