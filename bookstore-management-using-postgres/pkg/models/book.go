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
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookById(Id int) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("id = ?", Id).Find(&getbook)
	return &getbook, db
}
func DeleteBook(Id int) Book {
	var deletebook Book
	db.Where("id = ?", Id).Delete(deletebook)
	return deletebook
}
func UpdateBook(Id int, updatebook *Book) Book {
	db.Model(&updatebook).Where("id = ?", Id).Updates(Book{Name: updatebook.Name, Author: updatebook.Author, Price: updatebook.Price, Publication: updatebook.Publication})
	return *updatebook
}
