package models

import (
	"github.com/jishnu21/projectday3/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `gorm:"column:title" json:"title"`
	Author string `gorm:"column:author" json:"author"`
	Year   int    `gorm:"column:year" json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if db != nil {
		db.AutoMigrate(&Book{})
	} else {
		panic("Database connection is nil in models")
	}
}

// CreateBook inserts a new book into the database
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks returns all books from the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById returns a book by ID
func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	result := db.First(&book, id)
	return &book, result
}

// DeleteBook deletes a book by ID
func DeleteBook(id int64) Book {
	var book Book
	db.Delete(&book, id)
	return book
}
