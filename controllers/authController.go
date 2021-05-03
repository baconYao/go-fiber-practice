package controllers

import (
	"strconv"
	"time"

	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/baconYao/go-fiber-practice/util"
	"github.com/gofiber/fiber/v2"
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

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		RoleId: 1,	// 自動註冊為 admin. 2: editor. 3: viewer
	}
	// Hash password
	user.SetPassword(data["password"])

	// 使用 database 暴露的 DB，將 user 存進 DB
	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err :=c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	// 檢查 user 是否存在
	database.DB.Where("email = ?", []byte(data["email"])).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}
	// 檢查密碼是否相同
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	// get jwt token
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 建立 cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	// 取得名為 jwt 的 cookie
	cookie := c.Cookies("jwt")
	// Parse jwt token，並取的 user id
	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// 移除 cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",	// "" 表示要移除
		Expires: time.Now().Add(-time.Hour),	// 過去一小時
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}