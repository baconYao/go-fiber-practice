package controllers

import (
	"strconv"
	"time"

	"github.com/baconYao/go-fiber-practice/database"
	"github.com/baconYao/go-fiber-practice/models"
	"github.com/dgrijalva/jwt-go"
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
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),	// 1 day
	})

	token, err := claims.SignedString([]byte("secret"))

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

type Claims struct {
	// 取得 jwt.StandardClaims 的 fields (如下)
	jwt.StandardClaims
	// Audience  string `json:"aud,omitempty"`
	// ExpiresAt int64  `json:"exp,omitempty"`
	// Id        string `json:"jti,omitempty"`
	// IssuedAt  int64  `json:"iat,omitempty"`
	// Issuer    string `json:"iss,omitempty"`
	// NotBefore int64  `json:"nbf,omitempty"`
	// Subject   string `json:"sub,omitempty"`
}

func User(c *fiber.Ctx) error {
	// 取得名為 jwt 的 cookie
	cookie := c.Cookies("jwt")
	// 驗證並取得 jwt token
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// 取得 claims 內的 field 資訊
	claims := token.Claims.(*Claims)
	// 若沒有 .(*Claims)
	// claims := token.Claims
	// 則會回傳
	// {
	// 	"exp": 1619104831,
	// 	"iss": "1"
	// }
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	// TODO: 目前會回傳 user 的 password 欄位，雖然有 hash 過，但不該回傳給 user
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