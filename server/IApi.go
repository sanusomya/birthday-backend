package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Api interface{
	get(c context.Context, e echo.Context) error
	add(c context.Context, e echo.Context) error
	update(c context.Context, e echo.Context) error
	delete(c context.Context, e echo.Context) error
}