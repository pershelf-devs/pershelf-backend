package tablesModels

import (
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
}

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn"`
	Publisher     string    `json:"publisher"`
	PublishedYear int       `json:"published_year"`
	CoverImage    string    `json:"cover_image"`
	Genre         string    `json:"genre"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Reads         int       `json:"reads"`
	ImageBase64   string    `json:"image_base64"`
}

type Comment struct {
	ID           int       ` json:"id"`
	ReviewID     int       ` json:"review_id"`
	UserID       int       ` json:"user_id"`
	ParentCommID *int      ` json:"parent_comm_id"`
	CommentText  string    ` json:"comment_text"`
	CreatedAt    time.Time ` json:"created_at"`
	UpdatedAt    time.Time ` json:"updated_at"`
}

type Follow struct {
	ID         int       ` json:"id"`
	FollowerID int       ` json:"follower_id"`
	FollowedID int       ` json:"followed_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RefreshToken struct {
	ID        int       `json:"id"`
	UserID    int       ` json:"user_id"`
	Token     string    ` json:"token"`
	ExpiresAt time.Time ` json:"expires_at"`
	CreatedAt time.Time ` json:"created_at"`
}

type Review struct {
	ID         int       ` json:"id"`
	UserID     int       ` json:"user_id"`
	BookID     int       ` json:"book_id"`
	ReviewText string    ` json:"review_text"`
	CreatedAt  time.Time ` json:"created_at"`
	UpdatedAt  time.Time ` json:"updated_at"`
}

type ShelfBook struct {
	ID        int       ` json:"id"`
	ShelfID   int       ` json:"shelf_id"`
	BookID    int       ` json:"book_id"`
	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}

type UserBook struct {
	ID         int       ` json:"id"`
	UserID     int       ` json:"user_id"`
	BookID     int       ` json:"book_id"`
	Status     string    ` json:"status"`
	Rating     int       `json:"rating"`
	StartedAt  time.Time ` json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserShelf struct {
	ID        int       ` json:"id"`
	UserID    int       ` json:"user_id"`
	ShelfName string    ` json:"shelf_name"`
	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}
