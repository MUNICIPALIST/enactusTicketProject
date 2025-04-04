// internal/repository/postgres.go
package repository

import (
	"context"
	"ticket-service/internal/entity"
)

type Repository interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
	GetPrePosts(ctx context.Context) ([]entity.PrePost, error)
}
