package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// ShelfBook model
type ShelfBook struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	ShelfID   int       `gorm:"column:shelf_id;type:int(11);not null" json:"shelf_id"`
	BookID    int       `gorm:"column:book_id;type:int(11);not null" json:"book_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (ShelfBook) TableName() string {
	return "shelf_book"
}

// GetAllShelfBooks retrieves all shelf_book entries from the database
func GetAllShelfBooks() []ShelfBook {
	var shelfBooks []ShelfBook
	if err := globals.PershelfDB.Find(&shelfBooks).Error; err != nil {
		globals.Log("Error getting all shelf_books: ", err)
		return nil
	}
	return shelfBooks
}

// GetShelfBookByID retrieves a shelf_book by ID from the database
func GetShelfBookByID(id int) ShelfBook {
	var shelfBook ShelfBook
	if err := globals.PershelfDB.First(&shelfBook, id).Error; err != nil {
		globals.Log("Error getting shelf_book by ID: ", err)
		return ShelfBook{}
	}
	return shelfBook
}

// GetShelfBooksByShelfID retrieves all books in a specific shelf
func GetShelfBooksByShelfID(shelfID int) []ShelfBook {
	var shelfBooks []ShelfBook
	if err := globals.PershelfDB.Where("shelf_id = ?", shelfID).Find(&shelfBooks).Error; err != nil {
		globals.Log("Error getting shelf_books by shelf ID: ", err)
		return nil
	}
	return shelfBooks
}

// GetShelfBooksByBookID retrieves all shelves that contain a specific book
func GetShelfBooksByBookID(bookID int) []ShelfBook {
	var shelfBooks []ShelfBook
	if err := globals.PershelfDB.Where("book_id = ?", bookID).Find(&shelfBooks).Error; err != nil {
		globals.Log("Error getting shelf_books by book ID: ", err)
		return nil
	}
	return shelfBooks
}

// CreateShelfBook creates a new shelf_book entry in the database
func CreateShelfBook(shelfBook *ShelfBook) ShelfBook {
	if err := globals.PershelfDB.Create(&shelfBook).Error; err != nil {
		globals.Log("Error creating shelf_book: ", err)
		return ShelfBook{}
	}
	return *shelfBook
}

// UpdateShelfBook updates an existing shelf_book in the database
func UpdateShelfBook(shelfBook ShelfBook) ShelfBook {
	if err := globals.PershelfDB.Save(&shelfBook).Error; err != nil {
		globals.Log("Error updating shelf_book: ", err)
		return ShelfBook{}
	}
	return shelfBook
}

// DeleteShelfBook deletes a shelf_book entry from the database
func DeleteShelfBook(shelfBookID int) error {
	if err := globals.PershelfDB.Delete(&ShelfBook{}, shelfBookID).Error; err != nil {
		globals.Log("Error deleting shelf_book: ", err)
		return err
	}
	return nil
}
