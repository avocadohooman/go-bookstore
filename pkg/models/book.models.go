package models

import (
	"fmt"
	"github/avocadohooman/go-bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `"json:"name"`
	Author string `"json:"author"`
	Publication string `"json:"publication"`
}

func init() {
	config.ConnectDb()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	fmt.Println("Books", Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?",  Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

func (b *Book) UpdateBook(Id int64) (*Book) {
	db.Model(&b).Where("ID=?", Id).Updates(Book{Name: b.Name, Author: b.Author, Publication: b.Publication})
	return b
}
