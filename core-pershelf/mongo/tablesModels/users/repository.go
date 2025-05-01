package users

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
}

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) *UserRepositoryMongo {
	return &UserRepositoryMongo{collection: collection}
}

func (r *UserRepositoryMongo) CreateUser(user *User) error {
	log.Printf("DEBUG: Creating new user with username: %s", user.Username)
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("ERROR: Failed to create user: %v", err)
		return err
	}
	log.Printf("DEBUG: User created successfully with ID: %v", user.ID)
	return nil
}

func (r *UserRepositoryMongo) GetUserByID(id primitive.ObjectID) (*User, error) {
	log.Printf("DEBUG: Getting user by ID: %v", id)
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Printf("ERROR: Failed to get user by ID: %v", err)
		return nil, err
	}
	log.Printf("DEBUG: User found with ID: %v", id)
	return &user, nil
}

func (r *UserRepositoryMongo) GetUserByUsername(username string) (*User, error) {
	log.Printf("DEBUG: Getting user by username: %s", username)
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("ERROR: Failed to get user by username: %v", err)
		return nil, err
	}
	log.Printf("DEBUG: User found with username: %s", username)
	return &user, nil
}

func (r *UserRepositoryMongo) UpdateUser(user *User) error {
	log.Printf("DEBUG: Updating user with ID: %v", user.ID)
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	if err != nil {
		log.Printf("ERROR: Failed to update user: %v", err)
		return err
	}
	log.Printf("DEBUG: User updated successfully with ID: %v", user.ID)
	return nil
}

func (r *UserRepositoryMongo) DeleteUser(id primitive.ObjectID) error {
	log.Printf("DEBUG: Deleting user with ID: %v", id)
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Printf("ERROR: Failed to delete user: %v", err)
		return err
	}
	log.Printf("DEBUG: User deleted successfully with ID: %v", id)
	return nil
}
