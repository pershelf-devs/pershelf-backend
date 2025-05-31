package tablesModels

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

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id" validate:"required"`
	Title       string             `bson:"title" json:"title" validate:"required,min=1,max=1024"`
	Author      string             `bson:"author" json:"author" validate:"required,min=1,max=1024"`
	ISBN        string             `bson:"isbn" json:"isbn" validate:"required,min=10,max=13"`
	Description string             `bson:"description" json:"description" validate:"max=5000"`
	PageCount   int                `bson:"page_count" json:"page_count" validate:"required,min=1"`
	PublishedAt time.Time          `bson:"published_at" json:"published_at" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at" validate:"required"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at" validate:"required"`
	OwnerID     primitive.ObjectID `bson:"owner_id" json:"owner_id" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required,oneof=available borrowed reserved"`
	Category    string             `bson:"category" json:"category" validate:"required,min=1,max=64"`
	Language    string             `bson:"language" json:"language" validate:"required,min=2,max=3"`
}

type UserBook struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	BookID     primitive.ObjectID `bson:"book_id" json:"book_id"`
	Status     string             `bson:"status" json:"status"`
	Rating     int                `bson:"rating" json:"rating"`
	StartedAt  *time.Time         `bson:"started_at,omitempty" json:"started_at,omitempty"`
	FinishedAt *time.Time         `bson:"finished_at,omitempty" json:"finished_at,omitempty"`
}
