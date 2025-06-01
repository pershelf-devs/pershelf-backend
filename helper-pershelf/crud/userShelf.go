package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

// UserShelf model
type UserShelf struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	ShelfName string    `gorm:"column:shelf_name;type:varchar(255);not null" json:"shelf_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName sets the table name for GORM
func (UserShelf) TableName() string {
	return "user_shelves"
}

// GetAllUserShelfs retrieves all user_shelf entries from the database
func GetAllUserShelfs() []UserShelf {
	var userShelfs []UserShelf
	if err := globals.PershelfDB.Find(&userShelfs).Error; err != nil {
		globals.Log("Error getting all user_shelfs: ", err)
		return nil
	}
	return userShelfs
}

// GetUserShelfByID retrieves a user_shelf by ID from the database
func GetUserShelfByID(id int) UserShelf {
	var userShelf UserShelf
	if err := globals.PershelfDB.First(&userShelf, id).Error; err != nil {
		globals.Log("Error getting user_shelf by ID: ", err)
		return UserShelf{}
	}
	return userShelf
}

// GetUserShelvesByUserID retrieves a user_shelf by user ID from the database
func GetUserShelvesByUserID(userID int) []UserShelf {
	var userShelves []UserShelf
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&userShelves).Error; err != nil {
		globals.Log("Error getting user_shelves by user ID: ", err)
		return nil
	}
	return userShelves
}

// GetUserShelfsByUserID retrieves all shelves belonging to a specific user
func GetUserShelfsByUserID(userID int) []UserShelf {
	var userShelfs []UserShelf
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&userShelfs).Error; err != nil {
		globals.Log("Error getting user_shelfs by user ID: ", err)
		return nil
	}
	return userShelfs
}

// GetUserShelfByUserIDAndName retrieves a user_shelf by user ID and name from the database
func GetUserShelfByUserIDAndName(userID int, name string) UserShelf {
	var userShelf UserShelf
	if err := globals.PershelfDB.Where("user_id = ? AND shelf_name = ?", userID, name).First(&userShelf).Error; err != nil {
		globals.Log("Error getting user_shelf by user ID and name: ", err)
		return UserShelf{}
	}
	return userShelf
}

// CreateUserShelf creates a new user_shelf entry in the database
func CreateUserShelf(userShelf *UserShelf) UserShelf {
	if err := globals.PershelfDB.Create(&userShelf).Error; err != nil {
		globals.Log("Error creating user_shelf: ", err)
		return UserShelf{}
	}
	return *userShelf
}

// UpdateUserShelf updates an existing user_shelf in the database
func UpdateUserShelf(userShelf UserShelf) UserShelf {
	if err := globals.PershelfDB.Save(&userShelf).Error; err != nil {
		globals.Log("Error updating user_shelf: ", err)
		return UserShelf{}
	}
	return userShelf
}

// DeleteUserShelf deletes a user_shelf entry from the database
func DeleteUserShelf(userShelfID int) error {
	if err := globals.PershelfDB.Delete(&UserShelf{}, userShelfID).Error; err != nil {
		globals.Log("Error deleting user_shelf: ", err)
		return err
	}
	return nil
}
