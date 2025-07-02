package domain

import (
	"time"

	"github.com/google/uuid"
)

type FeedItem struct {
	PostID    uuid.UUID `json:"post_id"`
	UserID    uuid.UUID `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
	Score     float64   `json:"score"`
}

type Feed struct {
	UserID uuid.UUID  `json:"user_id"`
	Posts  []FeedItem `json:"posts"`
}
