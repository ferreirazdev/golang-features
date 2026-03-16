package domain

import "time"

type Task struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	Done        bool
}
