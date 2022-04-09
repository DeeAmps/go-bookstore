package models

import (
	"github.com/DeeAmps/go-bookstore-api/pkg/db"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() *gorm.DB {
	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	res := db.Create(&book)
	return res
}

func GetBooks() []Book {
	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(bookId int) Book {
	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	var book Book
	db.Where("ID=?", bookId).Find(&book)
	return book
}

func DeleteBook(bookId int) Book {
	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	var book Book
	db.Where("ID=?", bookId).Delete(&book)
	return book
}
