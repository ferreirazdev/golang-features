package usecase

import (
	"context"
	"taskapi/internal/domain"
)

type DeleteTaskUseCase struct {
	repository domain.TaskRepository
}

func NewDeleteTaskUseCase(repo domain.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		repository: repo,
	}
}

func (uc *DeleteTaskUseCase) Execute(ctx context.Context, id int64) error {
	return uc.repository.Delete(ctx, id)
}
