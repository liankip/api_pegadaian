package main

import (
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app, err := InitializeApplication()
	if err != nil {
		log.Fatal("Failed to initialize application: ", err)
	}

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
