package books

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookByID(id primitive.ObjectID) (*Book, error)
	GetBooksByOwnerID(ownerID primitive.ObjectID) ([]*Book, error)
	GetBooksByStatus(status string) ([]*Book, error)
	UpdateBook(book *Book) error
	DeleteBook(id primitive.ObjectID) error
}

type BookRepositoryMongo struct {
	collection *mongo.Collection
}

func NewBookRepositoryMongo(collection *mongo.Collection) *BookRepositoryMongo {
	return &BookRepositoryMongo{collection: collection}
}

func (r *BookRepositoryMongo) CreateBook(book *Book) error {
	log.Printf("DEBUG: Creating new book with title: %s", book.Title)
	_, err := r.collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Printf("ERROR: Failed to create book: %v", err)
		return err
	}
	log.Printf("DEBUG: Book created successfully with ID: %v", book.ID)
	return nil
}

func (r *BookRepositoryMongo) GetBookByID(id primitive.ObjectID) (*Book, error) {
	log.Printf("DEBUG: Getting book by ID: %v", id)
	var book Book
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		log.Printf("ERROR: Failed to get book by ID: %v", err)
		return nil, err
	}
	log.Printf("DEBUG: Book found with ID: %v", id)
	return &book, nil
}

func (r *BookRepositoryMongo) GetBooksByOwnerID(ownerID primitive.ObjectID) ([]*Book, error) {
	log.Printf("DEBUG: Getting books by owner ID: %v", ownerID)
	cursor, err := r.collection.Find(context.Background(), bson.M{"owner_id": ownerID})
	if err != nil {
		log.Printf("ERROR: Failed to get books by owner ID: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var books []*Book
	if err = cursor.All(context.Background(), &books); err != nil {
		log.Printf("ERROR: Failed to decode books: %v", err)
		return nil, err
	}
	log.Printf("DEBUG: Found %d books for owner ID: %v", len(books), ownerID)
	return books, nil
}

func (r *BookRepositoryMongo) GetBooksByStatus(status string) ([]*Book, error) {
	log.Printf("DEBUG: Getting books by status: %s", status)
	cursor, err := r.collection.Find(context.Background(), bson.M{"status": status})
	if err != nil {
		log.Printf("ERROR: Failed to get books by status: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var books []*Book
	if err = cursor.All(context.Background(), &books); err != nil {
		log.Printf("ERROR: Failed to decode books: %v", err)
		return nil, err
	}
	log.Printf("DEBUG: Found %d books with status: %s", len(books), status)
	return books, nil
}

func (r *BookRepositoryMongo) UpdateBook(book *Book) error {
	log.Printf("DEBUG: Updating book with ID: %v", book.ID)
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": book.ID},
		bson.M{"$set": book},
	)
	if err != nil {
		log.Printf("ERROR: Failed to update book: %v", err)
		return err
	}
	log.Printf("DEBUG: Book updated successfully with ID: %v", book.ID)
	return nil
}

func (r *BookRepositoryMongo) DeleteBook(id primitive.ObjectID) error {
	log.Printf("DEBUG: Deleting book with ID: %v", id)
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Printf("ERROR: Failed to delete book: %v", err)
		return err
	}
	log.Printf("DEBUG: Book deleted successfully with ID: %v", id)
	return nil
}
