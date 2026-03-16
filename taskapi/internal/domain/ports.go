package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task Task) (Task, error)
	List(ctx context.Context) ([]Task, error)
	Delete(ctx context.Context, id int64) error
}
