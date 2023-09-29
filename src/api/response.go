package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muchrief/go_pijar/src/model"
)

func ErrInternalServerResponse(err error, c echo.Context) error {
	return c.JSON(
		http.StatusInternalServerError,
		model.ResponeApi[interface{}]{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	)
}

func ErrUnprocessableEntityResponse(err error, c echo.Context) error {
	return c.JSON(
		http.StatusUnprocessableEntity,
		model.ResponeApi[interface{}]{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		},
	)
}

func ErrNotFoundResponse(err error, c echo.Context) error {
	return c.JSON(
		http.StatusNotFound,
		model.ResponeApi[interface{}]{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		},
	)
}

func ErrBadRequestResponse(err error, c echo.Context) error {
	return c.JSON(
		http.StatusBadRequest,
		model.ResponeApi[interface{}]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	)
}

func SuccessResponse(data interface{}, c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		model.ResponeApi[interface{}]{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		},
	)
}
