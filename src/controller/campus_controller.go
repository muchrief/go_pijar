package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"
)

func GetCampusInfo(c echo.Context) error {
	id := c.Param("id")
	crepo := repo.NewCampusRepo(database.DB)

	campus, err := crepo.GetCampus(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound(err, c)
		}

		return ErrInternalServer(err, c)
	}

	return SuccessResponse(campus, c)
}

func GetCampusFaculties(c echo.Context) error {
	id := c.Param("id")
	crepo := repo.NewCampusRepo(database.DB)

	campus, err := crepo.GetCampusFaculties(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound(err, c)
		}

		return ErrInternalServer(err, c)
	}

	return SuccessResponse(campus, c)
}
