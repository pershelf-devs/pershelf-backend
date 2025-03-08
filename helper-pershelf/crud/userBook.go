package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

type UserBook struct {
	ID         int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	UserID     int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	BookID     int       `gorm:"column:book_id;type:int(11);not null" json:"book_id"`
	Status     string    `gorm:"column:status;type:varchar(512);not null" json:"status"`
	Rating     int       `gorm:"column:rating;type:int(11);not null" json:"rating"`
	StartedAt  time.Time `gorm:"column:started_at;type:timestamp;not null" json:"started_at"`
	FinishedAt time.Time `gorm:"column:finished_at;type:timestamp;not null" json:"finished_at"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (UserBook) TableName() string {
	return "user_book"
}

// GetAllUserBooks gets all user books from the database
func GetAllUserBooks() []UserBook {
	var userBooks []UserBook
	if err := globals.PershelfDB.Find(&userBooks).Error; err != nil {
		globals.Log("Error getting all user books: ", err)
		return nil
	}
	return userBooks
}

// GetUserBookByID gets a user book entry by id from the database
func GetUserBookByID(id int) UserBook {
	var userBook UserBook
	if err := globals.PershelfDB.First(&userBook, id).Error; err != nil {
		globals.Log("Error getting user book by id: ", err)
		return userBook
	}
	return userBook
}

// GetUserBooksByUserID gets all books associated with a user ID
func GetUserBooksByUserID(userID int) []UserBook {
	var userBooks []UserBook
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&userBooks).Error; err != nil {
		globals.Log("Error getting user books by userID ", userID, ":", err)
		return nil
	}
	return userBooks
}

// GetUserBookByBookID gets all user book entries associated with a specific book ID
func GetUserBookByBookID(bookID int) []UserBook {
	var userBooks []UserBook
	if err := globals.PershelfDB.Where("book_id = ?", bookID).Find(&userBooks).Error; err != nil {
		globals.Log("Error getting user books by bookID ", bookID, ":", err)
		return nil
	}
	return userBooks
}

// CreateUserBook creates a new user book entry in the database
func CreateUserBook(userBook *UserBook) UserBook {
	if err := globals.PershelfDB.Create(userBook).Error; err != nil {
		globals.Log("Error creating user book: ", err)
	}
	return *userBook
}

// UpdateUserBook updates a user book entry in the database
func UpdateUserBook(userBook *UserBook) UserBook {
	if err := globals.PershelfDB.Save(userBook).Error; err != nil {
		globals.Log("Error updating user book: ", err)
	}
	return *userBook
}

// DeleteUserBook deletes a user book entry from the database
func DeleteUserBook(bookID int) error {
	if err := globals.PershelfDB.Delete(&UserBook{}, bookID).Error; err != nil {
		globals.Log("Error deleting user book: ", err)
		return err
	}
	return nil
}
