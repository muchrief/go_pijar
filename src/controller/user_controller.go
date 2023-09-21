package controller

import (
	"errors"
	"math"
	"net/http"

	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/auth"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	repo := repo.NewUserRepo(database.DB)
	var p model.Pagination
	err := c.Bind(&p)
	if err != nil {
		return ErrInternalServer(err, c)
	}

	if p.Sort == "" {
		p.Sort = "created_at"
	}

	users, err := repo.GetAllUser()
	if err != nil {
		return ErrInternalServer(err, c)
	}
	p.Total = int64(len(users))
	countPages := int(math.Ceil(float64(len(users)) / float64(p.Limit)))
	p.TotalPages = countPages

	data, err := repo.PaginateUser(&p)
	if err != nil {
		return ErrInternalServer(err, c)
	}
	p.Data = data

	resp := model.ResponeApi[model.Pagination]{
		Code:    0,
		Message: "Success",
		Data:    p,
	}

	return c.JSON(http.StatusOK, resp)
}

func AddUser(c echo.Context) error {
	user := new(model.AddUserPayload)
	err := c.Bind(user)
	if err != nil {
		return ErrInternalServer(err, c)
	}

	err = c.Validate(user)
	if err != nil {
		return ErrUnprocessableEntity(err, c)
	}

	repo := repo.NewUserRepo(database.DB)
	_, err = repo.GetUserByEmail(user.Email)
	if err == nil {
		return ErrBadRequest(
			errors.New("user has been added"),
			c,
		)
	}

	dbUser := &db_model.User{
		Username: user.Email,
		Password: user.Password,
		Role:     db_model.PUBLIC,
	}
	err = database.DB.Create(dbUser).Error
	if err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return ErrBadRequest(err, c)
		}

		return ErrInternalServer(err, c)
	}

	token, err := auth.GenerateAccessToken(&model.Auth{
		Id:       dbUser.Id.String(),
		Username: dbUser.Username,
		Role:     model.UserRole(dbUser.Role),
	})
	if err != nil {
		return ErrInternalServer(err, c)
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
		return ErrNotFound(err, c)
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
		return ErrInternalServer(err, c)
	}

	return c.JSON(http.StatusOK,
		model.ResponeApi[interface{}]{
			Code:    0,
			Message: "Success",
		},
	)
}
