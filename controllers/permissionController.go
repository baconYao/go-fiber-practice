package controllers

import (
	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
  var permissions []models.Permission

  database.DB.Find(&permissions)

  return c.JSON(permissions)
}