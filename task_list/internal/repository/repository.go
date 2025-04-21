package repository

import (
	"context"
	"database/sql"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
)

type CreateTaskDTO struct {
	TaskListID  int32
	UserID      int32
	Title       string
	Description sql.NullString
	Status      model.TaskStatus
}

type UpdateTaskDTO struct {
	TaskListID  int32
	UserID      int32
	Title       string
	Description sql.NullString
	Status      model.TaskStatus
}

type UpdateTaskListDTO struct {
	UserID      sql.NullInt32
	Title       sql.NullString
	Description sql.NullString
}

type PaginationOptions struct {
	Limit  int32
	Offset int32
}

type TaskListRepository interface {
	CreateTask(ctx context.Context, dto CreateTaskDTO) (int32, error)
	GetTaskByID(ctx context.Context, id int32) (*model.Task, error)
	UpdateTask(ctx context.Context, id int32, dto UpdateTaskDTO) error
	DeleteTask(ctx context.Context, id int32) error
	DeleteTaskBatch(ctx context.Context, idList []int32) error

	GetTaskListByID(ctx context.Context, id int32) (*model.TaskList, error)
	GetTasksByTaskList(ctx context.Context, taskListID int32, options PaginationOptions) ([]model.Task, error)
	GetTaskListsByUserID(ctx context.Context, userID int32, options PaginationOptions) ([]model.TaskList, error)
	DeleteTaskList(ctx context.Context, id int32) error
	UpdateTaskList(ctx context.Context, id int32, dto UpdateTaskListDTO) error
}
