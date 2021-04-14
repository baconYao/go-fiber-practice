package controllers

import (
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
	}
	user.LastName = "Yao"

	return c.JSON(user)
}