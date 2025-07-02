package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	MediaURL  string    `json:"media_url" db:"media_url"`
	MediaType string    `json:"media_type" db:"media_type"`
	Likes     int       `json:"likes" db:"likes"`
	Comments  int       `json:"comments" db:"comments"`
	Shares    int       `json:"shares" db:"shares"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	User      *User     `json:"user,omitempty"`
	IsLiked   bool      `json:"is_liked,omitempty"`
}

type Like struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	PostID    uuid.UUID `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Comment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PostID    uuid.UUID `json:"post_id" db:"post_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
	User *User `json:"user,omitempty"`
}
