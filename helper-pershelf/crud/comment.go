package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// Comment model
type Comment struct {
	ID           int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	ReviewID     int       `gorm:"column:review_id;type:int(11);not null" json:"review_id"`
	UserID       int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	ParentCommID *int      `gorm:"column:parent_comm_id;type:int(11);null" json:"parent_comm_id"`
	CommentText  string    `gorm:"column:comment_text;type:text;not null" json:"comment_text"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (Comment) TableName() string {
	return "comment"
}

// GetAllComments retrieves all comments from the database
func GetAllComments() []Comment {
	var comments []Comment
	if err := globals.PershelfDB.Find(&comments).Error; err != nil {
		globals.Log("Error getting all comments: ", err)
		return nil
	}
	return comments
}

// GetCommentByID retrieves a comment by ID from the database
func GetCommentByID(id int) Comment {
	var comment Comment
	if err := globals.PershelfDB.First(&comment, id).Error; err != nil {
		globals.Log("Error getting comment by ID: ", err)
		return Comment{}
	}
	return comment
}

// GetCommentsByReviewID retrieves all comments for a specific review
func GetCommentsByReviewID(reviewID int) []Comment {
	var comments []Comment
	if err := globals.PershelfDB.Where("review_id = ?", reviewID).Find(&comments).Error; err != nil {
		globals.Log("Error getting comments by review ID: ", err)
		return nil
	}
	return comments
}

// GetRepliesByCommentID retrieves all replies to a specific comment
func GetRepliesByCommentID(parentCommID int) []Comment {
	var replies []Comment
	if err := globals.PershelfDB.Where("parent_comm_id = ?", parentCommID).Find(&replies).Error; err != nil {
		globals.Log("Error getting replies by comment ID: ", err)
		return nil
	}
	return replies
}

// CreateComment creates a new comment in the database
func CreateComment(comment *Comment) Comment {
	if err := globals.PershelfDB.Create(&comment).Error; err != nil {
		globals.Log("Error creating comment: ", err)
		return Comment{}
	}
	return *comment
}

// UpdateComment updates an existing comment in the database
func UpdateComment(comment Comment) Comment {
	if err := globals.PershelfDB.Save(&comment).Error; err != nil {
		globals.Log("Error updating comment: ", err)
		return Comment{}
	}
	return comment
}

// DeleteComment deletes a comment from the database
func DeleteComment(commentID int) error {
	if err := globals.PershelfDB.Delete(&Comment{}, commentID).Error; err != nil {
		globals.Log("Error deleting comment: ", err)
		return err
	}
	return nil
}
