package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id" validate:"required"`
	Username    string             `bson:"username" json:"username" validate:"required,min=3,max=1024"`
	Password    string             `bson:"password" json:"password" validate:"required,min=6,max=128"`
	Email       string             `bson:"email" json:"email" validate:"email,max=1024"`
	Age         int                `bson:"age" json:"age" validate:"required,min=0,max=150"`
	Phone       string             `bson:"phone" json:"phone" validate:"max=16"`
	Description string             `bson:"description" json:"description" validate:"max=5000"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at" validate:"required"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at" validate:"required"`
	Name        string             `bson:"name" json:"name" validate:"required,max=1024"`
	Surname     string             `bson:"surname" json:"surname" validate:"required,max=64"`
}

