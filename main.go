package main

import (
	"go-gemini-api-sample/server"
)

func main() {
	webServer := server.NewEchoWebServer()
	webServer.Start()
}
