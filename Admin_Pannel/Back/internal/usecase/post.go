// internal/usecase/post.go
package usecase

import (
	"ticket-service/internal/entity"
	"ticket-service/internal/infrastructure/db"
)

type PostUseCase struct {
	repo *db.PostgresRepo
}

func NewPostUseCase(repo *db.PostgresRepo) *PostUseCase {
	return &PostUseCase{repo: repo}
}

func (uc *PostUseCase) GetPosts() ([]entity.Post, error) {
	return uc.repo.GetPosts()
}
