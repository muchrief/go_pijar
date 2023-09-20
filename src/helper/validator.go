package helper

import (
	"net/http"

	"github.com/muchrief/go_pijar/src/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			model.ResponeApi[interface{}]{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			})
	}

	return nil
}

func CreateCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}
