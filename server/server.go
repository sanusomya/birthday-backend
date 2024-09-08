package server

import (
	"os"
	"github.com/labstack/echo/v4"
)

func MyEchoServer() {
	var e = echo.New()
	registerEndpoints(e)

	port := os.Getenv("birthday_app_port")
	if port == ""{
		port = "8002"
	}
	e.Logger.Error(e.Start("localhost:"+port))
}
