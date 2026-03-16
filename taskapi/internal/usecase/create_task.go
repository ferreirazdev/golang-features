package usecase

import (
	"context"
	"errors"
	"taskapi/internal/domain"
	"time"
)

type CreateTaskInput struct {
	Title       string
	Description string
}

type CreateTaskUseCase struct {
	repository domain.TaskRepository
}

func NewCreateTaskUseCase(repo domain.TaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		repository: repo,
	}
}

func (uc *CreateTaskUseCase) Execute(ctx context.Context, input CreateTaskInput) (domain.Task, error) {
	if input.Title == "" {
		return domain.Task{}, errors.New("title is required")
	}

	task := domain.Task{
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	createdTask, err := uc.repository.Create(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return createdTask, nil
}
