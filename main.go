package main

import (
	"github.com/tkrajina/gofr"
	"your-module-name/handlers"
)

func main() {
	app := gofr.New()

	app.POST("/upload", handlers.UploadHandler)

	app.Start()
}
