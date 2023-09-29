package api

import (
	"math"
	"net/http"

	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/auth"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
	"github.com/muchrief/go_pijar/src/service"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	repo := repo.NewUserRepo(database.DB)
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	if p.Sort == "" {
		p.Sort = "created_at"
	}

	users, err := repo.GetAllUser()
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Total = int64(len(users))
	countPages := int(math.Ceil(float64(len(users)) / float64(p.GetLimit())))
	p.TotalPages = countPages

	data, err := repo.PaginateUser(&p)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}
	p.Data = data

	resp := model.ResponeApi[model.Pagination]{
		Code:    0,
		Message: "Success",
		Data:    p,
	}

	return SuccessResponse(resp, c)
}

func AddUser(c echo.Context) error {
	user := new(model.AddUserPayload)
	err := c.Bind(user)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	err = c.Validate(user)
	if err != nil {
		return ErrUnprocessableEntityResponse(err, c)
	}

	uService := service.NewUserService(database.DB)
	newUser, err := uService.CreateUser(user.Email, user.Password)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	token, err := auth.GenerateAccessToken(&model.Auth{
		Id:       newUser.Id.String(),
		Username: newUser.Username,
	})
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	resp := model.ResponeApi[model.AuthResp]{
		Code:    0,
		Message: "Success",
		Data: model.AuthResp{
			Token: token,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	repo := repo.NewUserRepo(database.DB)

	user, err := repo.GetUser(id)
	if err != nil {
		return ErrNotFoundResponse(err, c)
	}

	resp := model.ResponeApi[db_model.User]{
		Code:    0,
		Message: "Success",
		Data:    *user,
	}

	return c.JSON(http.StatusOK, resp)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	repo := repo.NewUserRepo(database.DB)
	err := repo.DeleteUser(id)
	if err != nil {
		return ErrInternalServerResponse(err, c)
	}

	return c.JSON(http.StatusOK,
		model.ResponeApi[interface{}]{
			Code:    0,
			Message: "Success",
		},
	)
}

func RegisterUserApi(app *echo.Echo) {
	g := app.Group("/users")

	g.POST("", AddUser)
	g.GET("", GetAllUser)
	g.GET("/:id", GetUser)
	g.DELETE("/:id", DeleteUser)
}
