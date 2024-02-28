package usecases

import (
	"go-gemini-api-sample/entities"
)

type IGeminiUsecase interface {
	GetGeminiResponse(question string) (*entities.Response, error)
}
