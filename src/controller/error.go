package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/src/model"
)

func ErrInternalServer(err error, c echo.Context) error {
	return c.JSON(
		http.StatusInternalServerError,
		model.ResponeApi[interface{}]{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	)
}

func ErrUnprocessableEntity(err error, c echo.Context) error {
	return c.JSON(
		http.StatusUnprocessableEntity,
		model.ResponeApi[interface{}]{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		},
	)
}

func ErrNotFound(err error, c echo.Context) error {
	return c.JSON(
		http.StatusNotFound,
		model.ResponeApi[interface{}]{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		},
	)
}

func ErrBadRequest(err error, c echo.Context) error {
	return c.JSON(
		http.StatusBadRequest,
		model.ResponeApi[interface{}]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	)
}
