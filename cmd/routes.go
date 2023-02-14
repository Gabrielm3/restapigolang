package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gabrielm3/restapigolang/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Put("/fact:id", handlers.UpdateFact)

	app.Delete("/fact:id", handlers.DeleteFact)
}
