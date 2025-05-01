package users

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
}

type UserServiceImpl struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (s *UserServiceImpl) CreateUser(user *User) error {
	log.Printf("DEBUG: Service layer - Creating user with username: %s", user.Username)

	// Zaman damgalarını ayarla
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Kullanıcı adı kontrolü
	existingUser, err := s.repository.GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		log.Printf("ERROR: Username already exists: %s", user.Username)
		return ErrUsernameExists
	}

	return s.repository.CreateUser(user)
}

func (s *UserServiceImpl) GetUserByID(id primitive.ObjectID) (*User, error) {
	log.Printf("DEBUG: Service layer - Getting user by ID: %v", id)
	return s.repository.GetUserByID(id)
}

func (s *UserServiceImpl) GetUserByUsername(username string) (*User, error) {
	log.Printf("DEBUG: Service layer - Getting user by username: %s", username)
	return s.repository.GetUserByUsername(username)
}

func (s *UserServiceImpl) UpdateUser(user *User) error {
	log.Printf("DEBUG: Service layer - Updating user with ID: %v", user.ID)

	// Güncelleme zamanını ayarla
	user.UpdatedAt = time.Now()

	// Kullanıcı adı değişikliği kontrolü
	if user.Username != "" {
		existingUser, err := s.repository.GetUserByUsername(user.Username)
		if err == nil && existingUser != nil && existingUser.ID != user.ID {
			log.Printf("ERROR: Username already exists: %s", user.Username)
			return ErrUsernameExists
		}
	}

	return s.repository.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id primitive.ObjectID) error {
	log.Printf("DEBUG: Service layer - Deleting user with ID: %v", id)
	return s.repository.DeleteUser(id)
}

// Hata tanımlamaları
var (
	ErrUsernameExists = &Error{Message: "username already exists"}
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
