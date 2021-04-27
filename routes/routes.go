package routes

import (
	"github.com/baconYao/go-fiber-practice/controllers"
	"github.com/baconYao/go-fiber-practice/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// 使用 auth middleware (以下的 api 都會有 auth 控管)
	app.Use(middlewares.IsAuthenticated)
	// 取得 user 資訊
	app.Get("/api/user", controllers.User)
	// 登出
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
}