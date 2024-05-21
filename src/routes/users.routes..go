package routes

import (
	"showdown-match-data/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/user", controllers.GetUserData)
}	