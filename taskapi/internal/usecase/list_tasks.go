package usecase

import (
	"context"
	"taskapi/internal/domain"
)

type ListTasksUseCase struct {
	repository domain.TaskRepository
}

func NewListTasksUseCase(repo domain.TaskRepository) *ListTasksUseCase {
	return &ListTasksUseCase{
		repository: repo,
	}
}

func (uc *ListTasksUseCase) Execute(ctx context.Context) ([]domain.Task, error) {
	return uc.repository.List(ctx)
}
