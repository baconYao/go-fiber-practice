package routes

import (
	"github.com/baconYao/go-fiber-practice/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}