package routes

import (
	"showdown-match-data/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func ReplaysRoutes(app *fiber.App) {
	app.Get("/replays", controllers.GetUserReplays)
	app.Get("/replays/users", controllers.GetUsersReplays)
	app.Get("/replays/format", controllers.GetFormatReplays)
	app.Get("/replays/user-format", controllers.GetUsersFormatReplays)
}