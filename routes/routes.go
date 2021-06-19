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

	// User 更新 profile
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	// 取得 user 資訊
	app.Get("/api/user", controllers.User)
	// 登出
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)

	app.Get("/api/products", controllers.AllProducts)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	app.Post("/api/upload", controllers.UploadImage)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/orders", controllers.AllOrders)
	app.Post("/api/export", controllers.Export)
	app.Get("/api/chart", controllers.Chart)
}