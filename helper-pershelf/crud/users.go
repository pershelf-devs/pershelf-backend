package crud

import (
	"time"

	"github.com/pershelf/pershelf/globals"
)

/*
MariaDB [pershelf]> describe users;
+-------------+---------------+------+-----+----------------------+----------------+
| Field       | Type          | Null | Key | Default              | Extra          |
+-------------+---------------+------+-----+----------------------+----------------+
| id          | int(11)       | NO   | PRI | NULL                 | auto_increment |
| username    | varchar(1024) | NO   | UNI | NULL                 |                |
| password    | varchar(128)  | YES  |     | NULL                 |                |
| email       | varchar(1024) | YES  |     | NULL                 |                |
| age         | int(11)       | NO   |     | NULL                 |                |
| phone       | varchar(16)   | YES  |     |                      |                |
| description | text          | YES  |     | NULL                 |                |
| created_at  | datetime(3)   | NO   |     | current_timestamp(3) |                |
| updated_at  | datetime(3)   | NO   |     | current_timestamp(3) |                |
| name        | text          | NO   |     | NULL                 |                |
| surname     | varchar(64)   | NO   |     | NULL                 |                |
+-------------+---------------+------+-----+----------------------+----------------+
*/

type User struct {
	ID          int       `gorm:"column:id;type:int(11);primaryKey;autoIncrement" json:"id"`
	Username    string    `gorm:"column:username;size:255;unique;not null" json:"username"`
	Password    string    `gorm:"column:password;size:128" json:"password"`
	Email       string    `gorm:"column:email;size:1024" json:"email"`
	Age         int       `gorm:"column:age;type:int(11);not null" json:"age"`
	Phone       string    `gorm:"column:phone;size:16;null;default:''" json:"phone"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	Name        string    `gorm:"column:name;type:text;not null" json:"name"`
	Surname     string    `gorm:"column:surname;type:varchar(64);not null" json:"surname"`
	ImageBase64 string    `gorm:"column:image_base64;type:longtext" json:"image_base64"`
}

func (User) TableName() string {
	return "users"
}

// GetAllUsers gets all users from the database
func GetAllUsers() []User {
	var users []User
	if err := globals.PershelfDB.Find(&users).Error; err != nil {
		globals.Log("Error getting all users: ", err)
		return nil
	}
	return users
}

// GetUserByID gets a user by id from the database
func GetUserByID(id int) User {
	var user User
	if err := globals.PershelfDB.First(&user, id).Error; err != nil {
		globals.Log("Error getting user by id: ", err)
		return User{}
	}
	return user
}

// GetUserByEmail gets a user by email from the database
func GetUserByEmail(email string) User {
	var user User
	if err := globals.PershelfDB.Where("email = ?", email).First(&user).Error; err != nil {
		globals.Log("Error getting user by email: ", err)
		return User{}
	}
	return user
}

// CreateUser creates a new user in the database
func CreateUser(user *User) User {
	if err := globals.PershelfDB.Create(&user).Error; err != nil {
		globals.Log("Error creating user: ", err)
		return User{}
	}
	return *user
}

// UpdateUser updates a user in the database
func UpdateUser(user User) User {
	if err := globals.PershelfDB.Save(&user).Error; err != nil {
		globals.Log("Error updating user: ", err)
		return User{}
	}
	return user
}

// DeleteUser deletes a user from the database
func DeleteUser(userID int) error {
	if err := globals.PershelfDB.Delete(&User{}, userID).Error; err != nil {
		globals.Log("Error deleting user: ", err)
		return err
	}
	return nil
}
