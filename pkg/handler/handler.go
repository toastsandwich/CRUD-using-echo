package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/CRUD-echo/pkg/model"
	"github.com/toastsandwich/CRUD-echo/pkg/utils"
)

func CreateUser(c echo.Context) error {
	utils.Lock.Lock()
	defer utils.Lock.Unlock()

	seq := utils.Seq
	u := &model.User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "binding error",
		})
	}

	utils.Users[seq] = u
	seq += 1
	return c.JSON(http.StatusAccepted, u)
}

func UpdateUser(c echo.Context) error {
	utils.Lock.Lock()
	defer utils.Lock.Unlock()

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "binding error",
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	utils.Users[id].Name = u.Name
	utils.Users[id].Email = u.Email
	return c.JSON(http.StatusAccepted, utils.Users[id])
}

func GetUsers(c echo.Context) error {
	utils.Lock.Lock()
	defer utils.Lock.Unlock()

	return c.JSON(http.StatusOK, utils.Users)
}

func GetUserByID(c echo.Context) error {
	utils.Lock.Lock()
	defer utils.Lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusFound, utils.Users[id])
}

func DeleteUserByID(c echo.Context) error {
	utils.Lock.Lock()
	defer utils.Lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))
	delete(utils.Users, id)
	return c.NoContent(http.StatusNoContent)
}
