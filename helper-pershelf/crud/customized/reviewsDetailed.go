package customized

import (
	"log"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/globals"
)

type DetailedReview struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	BookID      int       `json:"book_id"`
	ReviewText  string    `json:"review_text"`
	ReviewTitle string    `json:"review_title"`
	Rating      int       `json:"rating"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	User        crud.User `gorm:"foreignKey:UserID" json:"user"`
	Book        crud.Book `gorm:"foreignKey:BookID" json:"book"`
}

// GetDetailedReviewsByBookID gets detailed reviews by book ID
func GetDetailedReviewsByBookID(bookID int) ([]DetailedReview, error) {
	var detailedReviews []DetailedReview
	if err := globals.PershelfDB.Table("reviews").
		Where("book_id = ?", bookID).
		Preload("User").
		Preload("Book").
		Find(&detailedReviews).Error; err != nil {
		log.Printf("(Error): error getting detailed reviews by book ID: %v", err)
		return nil, err
	}
	return detailedReviews, nil
}
