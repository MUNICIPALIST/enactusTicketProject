package usecase

import (
	"ticket-service/internal/entity"
	"ticket-service/internal/infrastructure/db"
)

type PrePostUseCase struct {
	repo *db.PostgresRepo
}

func NewPrePostUseCase(repo *db.PostgresRepo) *PrePostUseCase {
	return &PrePostUseCase{repo: repo}
}

func (uc *PrePostUseCase) GetPrePosts() ([]entity.PrePost, error) {
	return uc.repo.GetPrePosts()
}
