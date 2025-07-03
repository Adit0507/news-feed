package ports

import (
	"context"

	"github.com/Adit0507/news-feed/internal/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (*domain.User, string, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, updates map[string]interface{}) error
	FollowUser(ctx context.Context, followerID, followingID uuid.UUID) error
	UnfollowUser(ctx context.Context, followerID, followingID uuid.UUID) error
	GetFollowers(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.User, error)
	GetFollowing(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, userID uuid.UUID, content, mediaURL, mediaType string) (*domain.Post, error)
	GetPost(ctx context.Context, postID uuid.UUID, userID *uuid.UUID) (*domain.Post, error)
	GetUserPosts(ctx context.Context, userID uuid.UUID, limit, offset int, viewerID *uuid.UUID) ([]domain.Post, error)
	UpdatePost(ctx context.Context, postID, userID uuid.UUID, content string) error
	DeletePost(ctx context.Context, postID, userID uuid.UUID) error
	LikePost(ctx context.Context, userID, postID uuid.UUID) error
	UnlikePost(ctx context.Context, userID, postID uuid.UUID) error
	AddComment(ctx context.Context, postID, userID uuid.UUID, content string) error
	GetComments(ctx context.Context, postID uuid.UUID, limit, offset int) ([]domain.Comment, error)
}

type FeedService interface {
	GetFeed(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.Post, error)
	RefreshFeed(ctx context.Context, userID uuid.UUID) error
	FanoutPost(ctx context.Context, post *domain.Post) error
}