package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

type RefreshToken struct {
	ID        int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	Token     string    `gorm:"column:token;type:varchar(512);not null" json:"token"`
	ExpiresAt time.Time `gorm:"column:expires_at;type:timestamp;not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}

// GetAllRefreshTokens gets all refreshTokens from the database
func GetAllRefreshTokens() []RefreshToken {
	var refreshTokens []RefreshToken
	if err := globals.PershelfDB.Find(&refreshTokens).Error; err != nil {
		globals.Log("Error getting all refresh tokens: ", err)
		return nil
	}
	return refreshTokens
}

// GetRefreshTokenByID gets a refresh token by id from the database
func GetRefreshTokenByID(id int) RefreshToken {
	var refreshToken RefreshToken
	if err := globals.PershelfDB.First(&refreshToken, id).Error; err != nil {
		globals.Log("Error getting refresh token by id: ", err)
		return RefreshToken{}
	}
	return refreshToken
}

// GetRefreshTokenByUserID gets a refresh token by UserID from the database
func GetRefreshTokenByUserID(userID int) RefreshToken {
	var refreshToken RefreshToken
	if err := globals.PershelfDB.Where("user_id = ?", userID).First(&refreshToken).Error; err != nil {
		globals.Log("Error getting refresh token by userID (", userID, "):", err)
		return RefreshToken{}
	}
	return refreshToken
}

// CreateRefreshToken creates a new refresh token in the database
func CreateRefreshToken(refreshToken *RefreshToken) RefreshToken {
	if err := globals.PershelfDB.Create(&refreshToken).Error; err != nil {
		globals.Log("Error creating refresh token: ", err)
		return RefreshToken{}
	}
	return *refreshToken
}

// UpdateRefreshToken updates a refresh token in the database
func UpdateRefreshToken(refreshToken RefreshToken) RefreshToken {
	if err := globals.PershelfDB.Save(&refreshToken).Error; err != nil {
		globals.Log("Error updating refresh token: ", err)
		return RefreshToken{}
	}
	return refreshToken
}

// DeleteRefreshToken deletes a refresh token from the database
func DeleteRefreshToken(tokenID int) error {
	if err := globals.PershelfDB.Delete(&RefreshToken{}, tokenID).Error; err != nil {
		globals.Log("Error deleting refresh token: ", err)
		return err
	}
	return nil
}
