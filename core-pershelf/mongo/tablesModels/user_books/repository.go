package user_books

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserBookRepository interface {
	CreateUserBook(userBook *UserBook) error
	GetUserBookByID(id primitive.ObjectID) (*UserBook, error)
	GetUserBooksByUserID(userID primitive.ObjectID) ([]*UserBook, error)
	UpdateUserBook(userBook *UserBook) error
	DeleteUserBook(id primitive.ObjectID) error
}

type UserBookRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserBookRepositoryMongo(collection *mongo.Collection) *UserBookRepositoryMongo {
	return &UserBookRepositoryMongo{collection: collection}
}

func (r *UserBookRepositoryMongo) CreateUserBook(userBook *UserBook) error {
	log.Printf("DEBUG: Creating new user_book for user: %v, book: %v", userBook.UserID, userBook.BookID)
	_, err := r.collection.InsertOne(context.Background(), userBook)
	if err != nil {
		log.Printf("ERROR: Failed to create user_book: %v", err)
		return err
	}
	return nil
}

func (r *UserBookRepositoryMongo) GetUserBookByID(id primitive.ObjectID) (*UserBook, error) {
	var userBook UserBook
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&userBook)
	if err != nil {
		return nil, err
	}
	return &userBook, nil
}

func (r *UserBookRepositoryMongo) GetUserBooksByUserID(userID primitive.ObjectID) ([]*UserBook, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var userBooks []*UserBook
	if err = cursor.All(context.Background(), &userBooks); err != nil {
		return nil, err
	}
	return userBooks, nil
}

func (r *UserBookRepositoryMongo) UpdateUserBook(userBook *UserBook) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userBook.ID},
		bson.M{"$set": userBook},
	)
	return err
}

func (r *UserBookRepositoryMongo) DeleteUserBook(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
