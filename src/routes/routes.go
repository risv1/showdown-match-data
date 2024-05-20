package routes

import (
	"showdown-match-data/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func MatchRoutes(app *fiber.App) {
	app.Get("/hello", controllers.Hello)
}	