package model

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusCreated    TaskStatus = "created"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCanceled   TaskStatus = "canceled"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID          int32
	TaskListID  int32
	Title       string
	Description string
	UserID      int32
	Status      TaskStatus
	CreatedAt   time.Time
}

type TaskList struct {
	ID          int32
	UserID      int32
	Title       string
	Description string
	CreatedAt   time.Time
}

type TaskListShare struct {
	TaskListID int32
	UserID     int32
}

type TaskShare struct {
	TaskID int32
	UserID int32
}
