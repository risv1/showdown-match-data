package main

import (
	"showdown-match-data/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.UserRoutes(app)
	routes.ReplaysRoutes(app)
	routes.DexRoutes(app)
	routes.LadderRoutes(app)

	app.Listen(":8000")
}