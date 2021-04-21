package main

import (
	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    database.Connect()
    app := fiber.New()
    // 為了要使前端能用 cookie，需要設定 AllowCredentials 為 true
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
    }))
    routes.Setup(app)

    app.Listen(":3000")
}