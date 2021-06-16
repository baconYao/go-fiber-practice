package controllers

import (
	"strconv"

	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	// Get query param 'page', its default value: 1
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}