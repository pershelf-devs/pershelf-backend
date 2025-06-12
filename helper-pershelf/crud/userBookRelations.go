package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

type UserBookRelation struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	BookID    int       `gorm:"column:book_id;type:int(11);not null" json:"book_id"`
	Like      bool      `gorm:"column:like;type:boolean;not null;default:false" json:"like"`
	Favorite  bool      `gorm:"column:favorite;type:boolean;not null;default:false" json:"favorite"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (UserBookRelation) TableName() string {
	return "user_book_relations"
}

// GetAllUserBookRelations gets all user book relations
func GetAllUserBookRelations() []UserBookRelation {
	userBookRelations := []UserBookRelation{}
	if err := globals.PershelfDB.Find(&userBookRelations).Error; err != nil {
		globals.Log("Error getting all user book relations", err)
		return nil
	}
	return userBookRelations
}

// GetUserBookRelationByID gets a user book relation by id
func GetUserBookRelationByID(id int) UserBookRelation {
	userBookRelation := UserBookRelation{}
	if err := globals.PershelfDB.Where("id = ?", id).First(&userBookRelation).Error; err != nil {
		globals.Log("Error getting user book relation by id", err)
		return UserBookRelation{}
	}
	return userBookRelation
}

// GetUserBookRelationsByUserID gets all user book relations by user id
func GetUserBookRelationsByUserID(userID int) []UserBookRelation {
	userBookRelations := []UserBookRelation{}
	if err := globals.PershelfDB.Where("user_id = ?", userID).Find(&userBookRelations).Error; err != nil {
		globals.Log("Error getting user book relations by user id", err)
		return nil
	}
	return userBookRelations
}

// GetUserBookRelationsByBookID gets all user book relations by book id
func GetUserBookRelationsByBookID(bookID int) []UserBookRelation {
	userBookRelations := []UserBookRelation{}
	if err := globals.PershelfDB.Where("book_id = ?", bookID).Find(&userBookRelations).Error; err != nil {
		globals.Log("Error getting user book relations by book id", err)
		return nil
	}
	return userBookRelations
}

// GetUserBookRelationByUserIDAndBookID gets a user book relation by user id and book id
func GetUserBookRelationByUserIDAndBookID(userID, bookID int) UserBookRelation {
	userBookRelation := UserBookRelation{}
	if err := globals.PershelfDB.Where("user_id = ? AND book_id = ?", userID, bookID).First(&userBookRelation).Error; err != nil {
		globals.Log("Error getting user book relation by user id and book id", err)
		return UserBookRelation{}
	}
	return userBookRelation
}

// CreateUserBookRelation creates a user book relation
func CreateUserBookRelation(userBookRelation *UserBookRelation) error {
	if err := globals.PershelfDB.Create(userBookRelation).Error; err != nil {
		globals.Log("Error creating user book relation", err)
		return err
	}
	return nil
}

// UpdateUserBookRelation updates a user book relation
func UpdateUserBookRelation(userBookRelation UserBookRelation) error {
	if err := globals.PershelfDB.Save(&userBookRelation).Error; err != nil {
		globals.Log("Error updating user book relation", err)
		return err
	}
	return nil
}

// DeleteUserBookRelationByID deletes a user book relation by id
func DeleteUserBookRelationByID(id int) error {
	if err := globals.PershelfDB.Delete(&UserBookRelation{}, id).Error; err != nil {
		globals.Log("Error deleting user book relation by id", err)
		return err
	}
	return nil
}

// DeleteUserBookRelationByUserID deletes a user book relation by user id
func DeleteUserBookRelationByUserID(userID int) error {
	if err := globals.PershelfDB.Delete(&UserBookRelation{}, "user_id = ?", userID).Error; err != nil {
		globals.Log("Error deleting user book relation by user id", err)
		return err
	}
	return nil
}

// DeleteUserBookRelationByBookID deletes a user book relation by book id
func DeleteUserBookRelationByBookID(bookID int) error {
	if err := globals.PershelfDB.Delete(&UserBookRelation{}, "book_id = ?", bookID).Error; err != nil {
		globals.Log("Error deleting user book relation by book id", err)
		return err
	}
	return nil
}

// DeleteUserBookRelationByUserIDAndBookID deletes a user book relation by user id and book id
func DeleteUserBookRelationByUserIDAndBookID(userID, bookID int) error {
	if err := globals.PershelfDB.Delete(&UserBookRelation{}, "user_id = ? AND book_id = ?", userID, bookID).Error; err != nil {
		globals.Log("Error deleting user book relation by user id and book id", err)
		return err
	}
	return nil
}
