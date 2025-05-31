package user_books

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBook struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	BookID     primitive.ObjectID `bson:"book_id" json:"book_id"`
	Status     string             `bson:"status" json:"status"`
	Rating     int                `bson:"rating" json:"rating"`
	StartedAt  *time.Time         `bson:"started_at,omitempty" json:"started_at,omitempty"`
	FinishedAt *time.Time         `bson:"finished_at,omitempty" json:"finished_at,omitempty"`
}
