package usecases

import (
	"github.com/google/generative-ai-go/genai"
)

type IGeminiUsecase interface {
	GetGeminiResponse(question string) (*genai.GenerateContentResponse, error)
}
