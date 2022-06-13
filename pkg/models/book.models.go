package models

import (
	"gorm.io/gorm"
	"github/avocadohooman/go-bookstore/pkg/config"
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

func (b* Book) CreateBook() *Book{
	found := db.Find(&b)
	if found == nil {
		db.Create(*b)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?",  Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

func (b* Book) UpdateBook(Id int64, updatedBook Book) (Book, *gorm.DB) {
	db := db.Model(&b).Updates(Book{Name: updatedBook.Name, Author: updatedBook.Author, Publication: updatedBook.Publication})
	return updatedBook, db
}
