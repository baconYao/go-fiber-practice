package database

import (
	"github.com/baconYao/go-fiber-practice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	database, err := gorm.Open(mysql.Open("root:rootroot@/go_admin"), &gorm.Config{})
    
    if err != nil {
        panic("Could not connect to the database")
    }

	database.AutoMigrate(&models.User{})
}