package controllers

import (
	"strconv"

	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	// Request Body
	// {
	// 		"name": "test2",
	// 		"permissions": ["1", "3"]
	// }

	// DTO : Data Transfer Object
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	// 取得 body 的 permissions 的值，是 list
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for idx, permissionId := range list {
		// 取得 request body 的 permissions list 的 id，並將其從 string 轉換成 int
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[idx] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role {
		Name: roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)
	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}

	result := database.DB.Preload("Permissions").Find(&role)
	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"message": "Role not found",
		})
	}
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	// DTO : Data Transfer Object
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	// 取得 body 的 permissions 的值，是 list
	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for idx, permissionId := range list {
		// 取得 request body 的 permissions list 的 id，並將其從 string 轉換成 int
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[idx] = models.Permission{
			Id: uint(id),
		}
	}
	// 根據 role_id，移除 role_permissions 表內相關的 records (permissions)，
	var result interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role {
		Id: uint(id),
		Name: roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}
