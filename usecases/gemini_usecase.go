package usecases

import (
	"context"
	"log"
	"os"

	"go-gemini-api-sample/entities"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiUsecase struct {}

func NewGeminiUsecase() IGeminiUsecase {
	return &GeminiUsecase{}
}

func (geminiUsecase *GeminiUsecase) GetGeminiResponse(question string) (*entities.Response, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Printf("Error creating genai client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	geminiRes, err := model.GenerateContent(ctx, genai.Text(question))

	if err != nil {
		return nil, err
	}

	response := entities.Response{
		Message: geminiRes.Candidates[0].Content.Parts[0],
	}

	return &response, nil
}
