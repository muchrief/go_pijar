package api

import (
	"errors"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
	"github.com/muchrief/go_pijar/src/service"
	"gorm.io/gorm"
)

func GetLectures(c echo.Context) error {
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

func LectureDetail(c echo.Context) error {
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

func LectureCourse(c echo.Context) error {
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

func AddLecture(c echo.Context) error {
	var payload model.LecturePayload
	err := c.Bind(&payload)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	err = c.Validate(payload)
	if err != nil {
		return ErrUnprocessableEntityResponse(err, c)
	}

	userRepo := repo.NewUserRepo(database.DB)
	_, err = userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, repo.ErrUserNotFound) {
			err := errors.New("email belum terdaftar")
			return ErrBadRequestResponse(err, c)
		}
		return ErrInternalServerResponse(err, nil)
	}

	lectureService := service.NewLectureService(database.DB)
	lecture := &db_model.Lecture{
		SchoolName:   payload.SchoolName,
		SupervisorId: payload.SupervisorId,
		Name:         payload.Name,
		Title:        payload.Title,
	}
	result, err := lectureService.AddLecture(lecture)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(result, c)
}

func UpdateLecture(c echo.Context) error {
	var payload model.LecturePayload
	err := c.Bind(&payload)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	err = c.Validate(payload)
	if err != nil {
		return ErrUnprocessableEntityResponse(err, c)
	}

	userRepo := repo.NewUserRepo(database.DB)
	_, err = userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, repo.ErrUserNotFound) {
			err := errors.New("email belum terdaftar")
			return ErrBadRequestResponse(err, c)
		}
		return ErrInternalServerResponse(err, nil)
	}

	lectureService := service.NewLectureService(database.DB)
	lecture := &db_model.Lecture{
		SchoolName:   payload.SchoolName,
		SupervisorId: payload.SupervisorId,
		Name:         payload.Name,
		Title:        payload.Title,
	}
	result, err := lectureService.UpdateLecture(lecture)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	return SuccessResponse(result, c)
}

func RegisterLectureApi(app *echo.Echo) {
	g := app.Group("/lectures")

	g.GET("", GetLectures)
	g.GET("/:id", LectureDetail)
	g.GET("/:id/courses", LectureCourse)

	g.POST("", AddLecture)
	g.PUT("", UpdateLecture)
}
