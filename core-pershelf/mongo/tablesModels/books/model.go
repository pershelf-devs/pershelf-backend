package books

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
