package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// Review model
type Review struct {
	ID          int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	UserID      int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	BookID      int       `gorm:"column:book_id;type:int(11);not null" json:"book_id"`
	ReviewTitle string    `gorm:"column:review_title;type:varchar(512);not null" json:"review_title"`
	ReviewText  string    `gorm:"column:review_text;type:text;not null" json:"review_text"`
	Rating      int       `gorm:"column:rating;type:int(11);not null" json:"rating"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (Review) TableName() string {
	return "reviews"
}

// GetAllReviews retrieves all reviews from the database
func GetAllReviews() []Review {
	var reviews []Review
	if err := globals.PershelfDB.Find(&reviews).Error; err != nil {
		globals.Log("Error getting all reviews: ", err)
		return nil
	}
	return reviews
}

// GetReviewByID retrieves a review by ID from the database
func GetReviewByID(id int) Review {
	var review Review
	if err := globals.PershelfDB.First(&review, id).Error; err != nil {
		globals.Log("Error getting review by ID: ", err)
		return Review{}
	}
	return review
}

// GetReviewsByUserID retrieves all reviews by a specific user ID
func GetReviewsByUserID(userID int) []Review {
	var reviews []Review
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		globals.Log("Error getting reviews by userID (", userID, "):", err)
		return nil
	}
	return reviews
}

// GetReviewsByBookID retrieves all reviews for a specific book ID
func GetReviewsByBookID(bookID int) []Review {
	var reviews []Review
	if err := globals.PershelfDB.Where("book_id = ?", bookID).Order("created_at DESC").Find(&reviews).Error; err != nil {
		globals.Log("Error getting reviews by bookID (", bookID, "):", err)
		return nil
	}
	return reviews
}

// CreateReview creates a new review in the database
func CreateReview(review *Review) Review {
	if err := globals.PershelfDB.Create(&review).Error; err != nil {
		globals.Log("Error creating review: ", err)
		return Review{}
	}
	return *review
}

// UpdateReview updates a review in the database
func UpdateReview(review Review) Review {
	if err := globals.PershelfDB.Save(&review).Error; err != nil {
		globals.Log("Error updating review: ", err)
		return Review{}
	}
	return review
}

// DeleteReview deletes a review from the database
func DeleteReview(reviewID int) error {
	if err := globals.PershelfDB.Delete(&Review{}, reviewID).Error; err != nil {
		globals.Log("Error deleting review: ", err)
		return err
	}
	return nil
}
