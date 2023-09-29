package api

import (
	"math"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
)

func getFaculties(c echo.Context) error {
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	if p.Sort == "" {
		p.Sort = "id"
	}

	repo := repo.NewFacultyRepo(database.DB)

	faculties, err := repo.Faculties(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Total = int64(len(faculties))
	countPages := int(math.Ceil(float64(len(faculties)) / float64(p.GetLimit())))
	p.TotalPages = countPages
	p.Data = faculties

	return SuccessResponse(p, c)
}

func getFacultySchools(c echo.Context) error {
	id := c.Param("id")
	repo := repo.NewFacultyRepo(database.DB)

	faculty, err := repo.FacultySchools(id)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(faculty, c)
}

func RegisterFacultyApi(app *echo.Echo) {
	g := app.Group("/faculties")

	g.GET("", getFaculties)
	g.GET("/:id", getFacultySchools)
}
