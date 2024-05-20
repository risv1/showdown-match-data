package main

import (
	"showdown-match-data/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.MatchRoutes(app)

	app.Listen(":8000")
}