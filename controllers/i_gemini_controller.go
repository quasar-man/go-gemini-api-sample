package controllers

import (
	"github.com/labstack/echo/v4"
)

type IGeminiController interface {
	PostQuestion(c echo.Context) error
}
