package api

import (
	"errors"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"
)

func getLectures(c echo.Context) error {
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	if p.Sort == "" {
		p.Sort = "id"
	}

	repo := repo.NewLectureRepo(database.DB)

	lectures, err := repo.GetLectures(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Total = int64(len(lectures))
	countPages := int(math.Ceil(float64(len(lectures)) / float64(p.GetLimit())))
	p.TotalPages = countPages
	p.Data = lectures

	return SuccessResponse(p, c)
}

func lectureDetail(c echo.Context) error {
	idParam := c.Param("id")

	repo := repo.NewLectureRepo(database.DB)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	lecture, err := repo.GetLectureCourse(int64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFoundResponse(err, c)
		}

		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(lecture, c)
}

func lectureCourse(c echo.Context) error {
	idParam := c.Param("id")

	repo := repo.NewLectureRepo(database.DB)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	lecture, err := repo.GetLectureCourse(int64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFoundResponse(err, c)
		}

		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(lecture, c)
}

func RegisterLectureApi(app *echo.Echo) {
	g := app.Group("/lectures")

	g.GET("", getLectures)
	g.GET("/:id", lectureDetail)
	g.GET("/:id/courses", lectureCourse)
}
