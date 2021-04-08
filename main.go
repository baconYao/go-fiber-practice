package main

import (
	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()
    app := fiber.New()
    routes.Setup(app)

    app.Listen(":3000")
}