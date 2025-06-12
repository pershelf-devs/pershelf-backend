package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// This will be deprecated (12.06.2025) => use UserBookRelation instead
type BookLike struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	BookID    int       `gorm:"column:book_id;type:int(11);not null" json:"book_id"`
	BookName  string    `gorm:"column:book_name;type:varchar(255);not null" json:"book_name"`
	UserID    int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	UserName  string    `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (BookLike) TableName() string {
	return "book_likes"
}

// GetBookLikesByBookID retrieves all book likes by book ID
func GetBookLikesByBookID(bookID int) []BookLike {
	var bookLikes []BookLike
	if err := globals.PershelfDB.Where("book_id = ?", bookID).Find(&bookLikes).Error; err != nil {
		globals.Log("Error getting book likes by book ID: ", err)
		return nil
	}
	return bookLikes
}

// GetBookLikesByUserID retrieves all book likes by user ID
func GetBookLikesByUserID(userID int) []BookLike {
	var bookLikes []BookLike
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&bookLikes).Error; err != nil {
		globals.Log("Error getting book likes by user ID: ", err)
		return nil
	}
	return bookLikes
}

// CreateBookLike creates a new book like
func CreateBookLike(bookLike *BookLike) error {
	if err := globals.PershelfDB.Create(&bookLike).Error; err != nil {
		globals.Log("Error creating book likes: ", err)
		return err
	}
	return nil
}

// UpdateBookLike updates a book like
func UpdateBookLike(bookLikes BookLike) error {
	if err := globals.PershelfDB.Save(&bookLikes).Error; err != nil {
		globals.Log("Error updating book likes: ", err)
		return err
	}
	return nil
}

// DeleteBookLikeByID deletes a book like by ID
func DeleteBookLikeByID(id int) error {
	if err := globals.PershelfDB.Delete(&BookLike{}, id).Error; err != nil {
		globals.Log("Error deleting book likes by ID: ", err)
		return err
	}
	return nil
}

// DeleteBookLikesByBookIDAndUserID deletes a book like by book ID and user ID
func DeleteBookLikesByBookIDAndUserID(bookID int, userID int) error {
	if err := globals.PershelfDB.Where("book_id = ? AND user_id = ?", bookID, userID).Delete(&BookLike{}).Error; err != nil {
		globals.Log("Error deleting book likes by book ID and user ID: ", err)
		return err
	}
	return nil
}
