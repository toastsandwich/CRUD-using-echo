package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/CRUD-echo/pkg/handler"
)

var Routes = func(e *echo.Echo) {
	e.GET("/api/allUsers", handler.GetUsers)
	e.GET("/api/user/:id", handler.GetUserByID)
	e.POST("/api/user", handler.CreateUser)
	e.PUT("/api/user/:id", handler.UpdateUser)
	e.DELETE("/api/user/:id", handler.DeleteUserByID)
}
