package tests

import (
	"context"
	"errors"
	"taskapi/internal/domain"
	"taskapi/internal/usecase"
	"testing"
)

type fakeTaskRepo struct {
	createCalled bool
	createArg    domain.Task
	nextID       int64
}

func (f *fakeTaskRepo) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	f.createCalled = true
	f.createArg = task
	task.ID = f.nextID
	f.nextID++
	return task, nil
}

func (f *fakeTaskRepo) List(ctx context.Context) ([]domain.Task, error) {
	return nil, nil
}

func (f *fakeTaskRepo) Delete(ctx context.Context, id int64) error {
	return nil
}

func TestCreateTaskUseCase_Execute_EmptyTitle_ReturnsError(t *testing.T) {
	ctx := context.Background()
	fake := &fakeTaskRepo{}
	uc := usecase.NewCreateTaskUseCase(fake)

	_, err := uc.Execute(ctx, usecase.CreateTaskInput{Title: "", Description: "desc"})

	if err == nil {
		t.Fatal("expected error for empty title")
	}
	if err.Error() != "title is required" {
		t.Errorf("expected 'title is required', got %q", err.Error())
	}
	if fake.createCalled {
		t.Error("repository Create should not be called when title is empty")
	}
}

func TestCreateTaskUseCase_Execute_ReturnsTaskWithID(t *testing.T) {
	ctx := context.Background()
	fake := &fakeTaskRepo{nextID: 1}
	uc := usecase.NewCreateTaskUseCase(fake)

	task, err := uc.Execute(ctx, usecase.CreateTaskInput{Title: "Buy milk", Description: "From the store"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.ID <= 0 {
		t.Errorf("expected task ID > 0, got %d", task.ID)
	}
	if task.Title != "Buy milk" {
		t.Errorf("expected Title 'Buy milk', got %q", task.Title)
	}
	if task.Description != "From the store" {
		t.Errorf("expected Description 'From the store', got %q", task.Description)
	}
	if task.Done {
		t.Error("expected Done false for new task")
	}
}

func TestCreateTaskUseCase_Execute_CallsRepoCreate(t *testing.T) {
	ctx := context.Background()
	fake := &fakeTaskRepo{nextID: 1}
	uc := usecase.NewCreateTaskUseCase(fake)

	_, err := uc.Execute(ctx, usecase.CreateTaskInput{Title: "Task", Description: "Desc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !fake.createCalled {
		t.Fatal("repository Create was not called")
	}
	if fake.createArg.Title != "Task" {
		t.Errorf("Create called with Title %q, want %q", fake.createArg.Title, "Task")
	}
	if fake.createArg.Description != "Desc" {
		t.Errorf("Create called with Description %q, want %q", fake.createArg.Description, "Desc")
	}
	if fake.createArg.Done {
		t.Error("Create should be called with Done false")
	}
	if fake.createArg.ID != 0 {
		t.Errorf("Create should be called with ID 0 (assigning ID is repo's job), got %d", fake.createArg.ID)
	}
}

func TestCreateTaskUseCase_Execute_PropagatesRepoError(t *testing.T) {
	ctx := context.Background()
	repoErr := errors.New("db unavailable")
	errRepo := &errorFakeTaskRepo{err: repoErr}
	uc := usecase.NewCreateTaskUseCase(errRepo)

	_, err := uc.Execute(ctx, usecase.CreateTaskInput{Title: "OK", Description: ""})

	if err != repoErr {
		t.Errorf("expected repo error to be propagated, got %v", err)
	}
}

// errorFakeTaskRepo only implements Create and returns a fixed error (for PropagatesRepoError test).
type errorFakeTaskRepo struct{ err error }

func (e *errorFakeTaskRepo) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	return domain.Task{}, e.err
}
func (e *errorFakeTaskRepo) List(ctx context.Context) ([]domain.Task, error) {
	return nil, nil
}
func (e *errorFakeTaskRepo) Delete(ctx context.Context, id int64) error {
	return nil
}
