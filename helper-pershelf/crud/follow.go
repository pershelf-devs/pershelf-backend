package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// Follow model
type Follow struct {
	ID         int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	FollowerID int       `gorm:"column:follower_id;type:int(11);not null" json:"follower_id"`
	FollowedID int       `gorm:"column:followed_id;type:int(11);not null" json:"followed_id"`
	Status     string    `gorm:"column:status;type:varchar(20);not null;default:'pending'" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (Follow) TableName() string {
	return "follow"
}

// GetAllFollows retrieves all follow relationships from the database
func GetAllFollows() []Follow {
	var follows []Follow
	if err := globals.PershelfDB.Find(&follows).Error; err != nil {
		globals.Log("Error getting all follow relationships: ", err)
		return nil
	}
	return follows
}

// GetFollowByID retrieves a follow relationship by ID
func GetFollowByID(id int) Follow {
	var follow Follow
	if err := globals.PershelfDB.First(&follow, id).Error; err != nil {
		globals.Log("Error getting follow by ID: ", err)
		return Follow{}
	}
	return follow
}

// GetFollowers retrieves all followers of a specific user
func GetFollowers(userID int) []Follow {
	var followers []Follow
	if err := globals.PershelfDB.Where("followed_id = ?", userID).Find(&followers).Error; err != nil {
		globals.Log("Error getting followers: ", err)
		return nil
	}
	return followers
}

// GetFollowing retrieves all users that a specific user follows
func GetFollowing(userID int) []Follow {
	var following []Follow
	if err := globals.PershelfDB.Where("follower_id = ?", userID).Find(&following).Error; err != nil {
		globals.Log("Error getting following list: ", err)
		return nil
	}
	return following
}

// CreateFollow creates a new follow relationship
func CreateFollow(follow *Follow) Follow {
	if err := globals.PershelfDB.Create(&follow).Error; err != nil {
		globals.Log("Error creating follow relationship: ", err)
		return Follow{}
	}
	return *follow
}

// UpdateFollow updates a follow relationship
func UpdateFollow(follow Follow) Follow {
	if err := globals.PershelfDB.Save(&follow).Error; err != nil {
		globals.Log("Error updating follow relationship: ", err)
		return Follow{}
	}
	return follow
}

// DeleteFollow deletes a follow relationship
func DeleteFollow(followID int) error {
	if err := globals.PershelfDB.Delete(&Follow{}, followID).Error; err != nil {
		globals.Log("Error deleting follow relationship: ", err)
		return err
	}
	return nil
}
