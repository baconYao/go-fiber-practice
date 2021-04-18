package controllers

import (
	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err :=c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password doesn't match",
		})
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	
	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		Password: hashedPassword,
	}

	// 使用 database 暴露的 DB，將 user 存進 DB
	database.DB.Create(&user)

	return c.JSON(user)
}