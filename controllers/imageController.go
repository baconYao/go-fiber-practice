package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
  form, err := c.MultipartForm()

  if err != nil {
    return err
  }

  // form 有個 image 名稱的欄位
  files := form.File["image"]

  fileName := ""
  for _, file := range files {
    fileName = file.Filename
    fmt.Println(fileName)
    if err := c.SaveFile(file, "./uploads/" + fileName); err != nil {
      return err
    }
  }
  return c.JSON(fiber.Map{
    "url": "http://localhost:3000/api/uploads/" + fileName,
  })
}