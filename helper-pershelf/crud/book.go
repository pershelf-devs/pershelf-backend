package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// Book model
type Book struct {
	ID            int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	Title         string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Author        string    `gorm:"column:author;type:varchar(255);not null" json:"author"`
	ISBN          string    `gorm:"column:isbn;type:varchar(20);unique;not null" json:"isbn"`
	Publisher     string    `gorm:"column:publisher;type:varchar(255);not null" json:"publisher"`
	PublishedYear int       `gorm:"column:published_year;type:int(4);not null" json:"published_year"`
	CoverImage    string    `gorm:"column:cover_image;type:varchar(512);null" json:"cover_image"`
	Genre         string    `gorm:"column:genre;type:varchar(100);not null" json:"genre"`
	Description   string    `gorm:"column:description;type:text;null" json:"description"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (Book) TableName() string {
	return "book"
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
	var books []Book
	if err := globals.PershelfDB.Find(&books).Error; err != nil {
		globals.Log("Error getting all books: ", err)
		return nil
	}
	return books
}

// GetBookByID retrieves a book by ID from the database
func GetBookByID(id int) Book {
	var book Book
	if err := globals.PershelfDB.First(&book, id).Error; err != nil {
		globals.Log("Error getting book by ID: ", err)
		return Book{}
	}
	return book
}

// GetBookByISBN retrieves a book by ISBN from the database
func GetBookByISBN(isbn string) Book {
	var book Book
	if err := globals.PershelfDB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		globals.Log("Error getting book by ISBN (", isbn, "):", err)
		return Book{}
	}
	return book
}

// CreateBook creates a new book in the database
func CreateBook(book *Book) Book {
	if err := globals.PershelfDB.Create(&book).Error; err != nil {
		globals.Log("Error creating book: ", err)
		return Book{}
	}
	return *book
}

// UpdateBook updates an existing book in the database
func UpdateBook(book Book) Book {
	if err := globals.PershelfDB.Save(&book).Error; err != nil {
		globals.Log("Error updating book: ", err)
		return Book{}
	}
	return book
}

// DeleteBook deletes a book from the database
func DeleteBook(bookID int) error {
	if err := globals.PershelfDB.Delete(&Book{}, bookID).Error; err != nil {
		globals.Log("Error deleting book: ", err)
		return err
	}
	return nil
}
