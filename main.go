package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(mysql.Open("root:rootroot@/go_admin"), &gorm.Config{})
    
    if err != nil {
        panic("Could not connect to the database")
    }

    fmt.Println(db)

    app := fiber.New()

    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}