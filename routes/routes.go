package routes

import (
	"github.com/baconYao/go-fiber-practice/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	// 取得 user 資訊
	app.Get("/api/user", controllers.User)
}