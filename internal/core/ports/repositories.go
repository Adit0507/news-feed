package ports

import (
	"context"

	"github.com/Adit0507/news-feed/internal/core/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
    GetByEmail(ctx context.Context, email string) (*domain.User, error)
    GetByUsername(ctx context.Context, username string) (*domain.User, error)
    Update(ctx context.Context, user *domain.User) error
    Follow(ctx context.Context, followerID, followingID uuid.UUID) error
    Unfollow(ctx context.Context, followerID, followingID uuid.UUID) error
    GetFollowers(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.User, error)
    GetFollowing(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.User, error)
    IsFollowing(ctx context.Context, followerID, followingID uuid.UUID) (bool, error)
}

type PostRepository interface {
    Create(ctx context.Context, post *domain.Post) error
    GetByID(ctx context.Context, id uuid.UUID) (*domain.Post, error)
    GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.Post, error)
    Update(ctx context.Context, post *domain.Post) error
    Delete(ctx context.Context, id uuid.UUID) error
    LikePost(ctx context.Context, userID, postID uuid.UUID) error
    UnlikePost(ctx context.Context, userID, postID uuid.UUID) error
    IsLiked(ctx context.Context, userID, postID uuid.UUID) (bool, error)
    GetLikes(ctx context.Context, postID uuid.UUID, limit, offset int) ([]domain.User, error)
    AddComment(ctx context.Context, comment *domain.Comment) error
    GetComments(ctx context.Context, postID uuid.UUID, limit, offset int) ([]domain.Comment, error)
    GetFeedPosts(ctx context.Context, userIDs []uuid.UUID, limit, offset int) ([]domain.Post, error)
}

type FeedRepository interface {
    CacheFeed(ctx context.Context, userID uuid.UUID, posts []domain.FeedItem) error
    GetCachedFeed(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.FeedItem, error)
    InvalidateFeed(ctx context.Context, userID uuid.UUID) error
    AddToFanoutFeeds(ctx context.Context, post *domain.Post, followerIDs []uuid.UUID) error
}
