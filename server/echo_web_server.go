package server

import (
	"go-gemini-api-sample/controllers"

	"github.com/labstack/echo/v4"
)

type EchoWebServer struct {}

func NewEchoWebServer() IServer {
	return &EchoWebServer{}
}

func (echoWebServer *EchoWebServer) Start() {
	e := echo.New()

	geminiController := controllers.NewGeminiController()
	e.POST("/post_question", geminiController.PostQuestion)
	e.Logger.Fatal(e.Start(":1323"))
}
