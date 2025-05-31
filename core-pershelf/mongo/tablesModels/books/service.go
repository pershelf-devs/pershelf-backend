package books

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService interface {
	CreateBook(book *Book) error
	GetBookByID(id primitive.ObjectID) (*Book, error)
	GetBooksByOwnerID(ownerID primitive.ObjectID) ([]*Book, error)
	GetBooksByStatus(status string) ([]*Book, error)
	UpdateBook(book *Book) error
	DeleteBook(id primitive.ObjectID) error
}

type BookServiceImpl struct {
	repository BookRepository
}

func NewBookService(repository BookRepository) *BookServiceImpl {
	return &BookServiceImpl{repository: repository}
}

func (s *BookServiceImpl) CreateBook(book *Book) error {
	log.Printf("DEBUG: Service layer - Creating book with title: %s", book.Title)

	// Zaman damgalarını ayarla
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	return s.repository.CreateBook(book)
}

func (s *BookServiceImpl) GetBookByID(id primitive.ObjectID) (*Book, error) {
	log.Printf("DEBUG: Service layer - Getting book by ID: %v", id)
	return s.repository.GetBookByID(id)
}

func (s *BookServiceImpl) GetBooksByOwnerID(ownerID primitive.ObjectID) ([]*Book, error) {
	log.Printf("DEBUG: Service layer - Getting books by owner ID: %v", ownerID)
	return s.repository.GetBooksByOwnerID(ownerID)
}

func (s *BookServiceImpl) GetBooksByStatus(status string) ([]*Book, error) {
	log.Printf("DEBUG: Service layer - Getting books by status: %s", status)
	return s.repository.GetBooksByStatus(status)
}

func (s *BookServiceImpl) UpdateBook(book *Book) error {
	log.Printf("DEBUG: Service layer - Updating book with ID: %v", book.ID)

	// Güncelleme zamanını ayarla
	book.UpdatedAt = time.Now()

	return s.repository.UpdateBook(book)
}

func (s *BookServiceImpl) DeleteBook(id primitive.ObjectID) error {
	log.Printf("DEBUG: Service layer - Deleting book with ID: %v", id)
	return s.repository.DeleteBook(id)
}
