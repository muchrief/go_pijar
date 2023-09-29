package api

import (
	"errors"
	"math"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"
)

func GetCampus(c echo.Context) error {
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	if p.Sort == "" {
		p.Sort = "id"
	}

	repo := repo.NewCampusRepo(database.DB)

	campus, err := repo.Campus(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Total = int64(len(campus))
	countPages := int(math.Ceil(float64(len(campus)) / float64(p.GetLimit())))
	p.TotalPages = countPages
	p.Data = campus

	return SuccessResponse(p, c)

}

func GetCampusDetail(c echo.Context) error {
	id := c.Param("id")
	crepo := repo.NewCampusRepo(database.DB)

	campus, err := crepo.GetCampusDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFoundResponse(err, c)
		}

		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(campus, c)
}

func GetCampusFaculties(c echo.Context) error {
	id := c.Param("id")
	crepo := repo.NewCampusRepo(database.DB)

	campus, err := crepo.GetCampusFaculties(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFoundResponse(err, c)
		}

		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(campus, c)
}

func GetCampusSchools(c echo.Context) error {
	id := c.Param("id")
	crepo := repo.NewCampusRepo(database.DB)

	campus, err := crepo.GetCampusSchools(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFoundResponse(err, c)
		}

		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(campus, c)
}

func RegisterCampusApi(app *echo.Echo) {
	g := app.Group("/campus")

	g.GET("", GetCampus)
	g.GET("/:id", GetCampusDetail)
	g.GET("/:id/faculties", GetCampusFaculties)
	g.GET("/:id/schools", GetCampusSchools)
}
