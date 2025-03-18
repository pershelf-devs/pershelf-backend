package tablesModels

import "time"

// This file will contain database table models
type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
}
