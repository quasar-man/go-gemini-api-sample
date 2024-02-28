package usecases

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiUsecase struct {}

func NewGeminiUsecase() IGeminiUsecase {
	return &GeminiUsecase{}
}

func (geminiUsecase *GeminiUsecase) GetGeminiResponse(question string) (interface{}, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Printf("Error creating genai client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	response, err := model.GenerateContent(ctx, genai.Text(question))

	if err != nil {
		return nil, err
	}

	return response, nil
}
