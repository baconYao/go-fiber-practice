package database

import (
	"github.com/baconYao/go-fiber-practice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 暴露 DB 給其他檔案使用
var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:rootroot@/go_admin"), &gorm.Config{})
    
    if err != nil {
        panic("Could not connect to the database")
    }

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}