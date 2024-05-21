package routes

import (
	"github.com/gofiber/fiber/v2"
	"showdown-match-data/src/controllers"
)

func LadderRoutes(app *fiber.App) {
	app.Get("/ladder", controllers.GetFormatLadderData)
}
