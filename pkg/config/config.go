package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var App *echo.Echo

func BootUp() {
	App = echo.New()
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())
}
