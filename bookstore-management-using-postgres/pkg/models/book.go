package models

import (
	"github.com/Patil-Kranti/Learning-GO/tree/main/bookstore-management-using-postgress/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `gorm:"json:author"`
	Price       int    `gorm:"json:price"`
	Publication string `gorm:"json:publication"`
}

func Init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() Book {
	db.Create(&b)
	return *b
}
