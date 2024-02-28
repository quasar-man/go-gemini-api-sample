package controllers

import (
	"go-gemini-api-sample/usecases"

	"github.com/labstack/echo/v4"
)

type GeminiController struct {}

func NewGeminiController() IGeminiController {
	return &GeminiController{}
}

func (geminiController *GeminiController) PostQuestion(c echo.Context) error {
	geminiUsecase := usecases.NewGeminiUsecase()
	response, err := geminiUsecase.GetGeminiResponse(c.FormValue("question"))

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, response)
}
