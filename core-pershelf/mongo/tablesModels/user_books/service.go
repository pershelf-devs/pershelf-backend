package user_books

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBookService interface {
	CreateUserBook(userBook *UserBook) error
	GetUserBookByID(id primitive.ObjectID) (*UserBook, error)
	GetUserBooksByUserID(userID primitive.ObjectID) ([]*UserBook, error)
	UpdateUserBook(userBook *UserBook) error
	DeleteUserBook(id primitive.ObjectID) error
}

type UserBookServiceImpl struct {
	repository UserBookRepository
}

func NewUserBookService(repository UserBookRepository) *UserBookServiceImpl {
	return &UserBookServiceImpl{repository: repository}
}

func (s *UserBookServiceImpl) CreateUserBook(userBook *UserBook) error {
	log.Printf("DEBUG: Service layer - Creating user_book for user: %v, book: %v", userBook.UserID, userBook.BookID)
	userBook.ID = primitive.NewObjectID()
	if userBook.StartedAt == nil {
		now := time.Now()
		userBook.StartedAt = &now
	}
	return s.repository.CreateUserBook(userBook)
}

func (s *UserBookServiceImpl) GetUserBookByID(id primitive.ObjectID) (*UserBook, error) {
	return s.repository.GetUserBookByID(id)
}

func (s *UserBookServiceImpl) GetUserBooksByUserID(userID primitive.ObjectID) ([]*UserBook, error) {
	return s.repository.GetUserBooksByUserID(userID)
}

func (s *UserBookServiceImpl) UpdateUserBook(userBook *UserBook) error {
	return s.repository.UpdateUserBook(userBook)
}

func (s *UserBookServiceImpl) DeleteUserBook(id primitive.ObjectID) error {
	return s.repository.DeleteUserBook(id)
}
