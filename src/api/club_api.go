package api

import (
	"math"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
)

func Clubs(c echo.Context) error {
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	if p.Sort == "" {
		p.Sort = "id"
	}

	repo := repo.NewClubRepo(database.DB)

	clubs, err := repo.Clubs(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Total = int64(len(clubs))
	countPages := int(math.Ceil(float64(len(clubs)) / float64(p.GetLimit())))
	p.TotalPages = countPages
	p.Data = clubs

	return SuccessResponse(p, c)
}

func Club(c echo.Context) error {
	id := c.Param("id")
	repo := repo.NewClubRepo(database.DB)

	club, err := repo.Club(id)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(club, c)
}

func RegisterClubApi(app *echo.Echo) {
	g := app.Group("/club")

	g.GET("", Clubs)
	g.GET("/:id", Club)
}
