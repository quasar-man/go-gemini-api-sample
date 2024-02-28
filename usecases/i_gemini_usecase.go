package usecases

type IGeminiUsecase interface {
	GetGeminiResponse(question string) (interface{}, error)
}
