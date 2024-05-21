package routes

import (
	"showdown-match-data/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func DexRoutes(app *fiber.App) {
	app.Get("/dex", controllers.GetDexData)
	app.Get("/moves", controllers.GetMovesData)
}