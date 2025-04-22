package service

import (
	"context"
	"database/sql"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository"
	"log/slog"
)

func toNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

type TaskListService struct {
	r      repository.TaskListRepository
	logger *slog.Logger
}

func (t *TaskListService) GetTaskList(ctx context.Context, id int32) (*model.TaskList, error) {
	list, err := t.r.GetTaskListByID(ctx, id)
	if err != nil {
		t.logger.Error("GetTaskListByID", "error", err)
		return nil, err
	}

	return list, err
}

type CreateTaskListOptions struct {
	UserID      int32
	Title       string
	Description *string
}

func (t *TaskListService) CreateTaskList(ctx context.Context, data CreateTaskListOptions) (int32, error) {
	opts := repository.CreateTaskListDTO{
		UserID:      data.UserID,
		Title:       data.Title,
		Description: toNullString(data.Description),
	}

	id, err := t.r.CreateTaskList(ctx, opts)
	if err != nil {
		t.logger.Error("CreateTaskList", "error", err)
		return 0, err
	}

	return id, err
}

func (t *TaskListService) DeleteTaskList(ctx context.Context, id int32) error {
	err := t.r.DeleteTaskList(ctx, id)
	if err != nil {
		t.logger.Error("DeleteTaskList", "error", err)
		return err
	}

	return nil
}

type UpdateTaskListOptions struct {
	UserID      *int32
	Title       *string
	Description *string
}

func (t *TaskListService) UpdateTaskList(ctx context.Context, id int32, data UpdateTaskListOptions) error {
	opts := repository.UpdateTaskListDTO{
		UserID:      toNullInt32(data.UserID),
		Title:       toNullString(data.Title),
		Description: toNullString(data.Description),
	}

	err := t.r.UpdateTaskList(ctx, id, opts)
	if err != nil {
		t.logger.Error("UpdateTaskList", "error", err)
		return err
	}

	return nil
}

type GetTaskListsOptions struct {
	Limit  int32
	Offset int32
}

func (t *TaskListService) GetTaskListsByUserID(ctx context.Context, id int32, options GetTaskListsOptions) ([]model.TaskList, error) {
	l, err := t.r.GetTaskListsByUserID(ctx, id, repository.PaginationOptions{
		Limit:  options.Limit,
		Offset: options.Offset,
	})
	if err != nil {
		t.logger.Error("GetTaskListsByUserID", "error", err)
		return nil, err
	}

	return l, err
}

type CreateTaskOptions struct {
}

func (t *TaskListService) CreateTask(ctx context.Context, data repository.CreateTaskDTO) (int32, error) {
	return 0, nil
}
