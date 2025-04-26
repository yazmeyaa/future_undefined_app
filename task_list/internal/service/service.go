package service

import (
	"context"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository"
	"log/slog"
)

type TaskListService struct {
	r      repository.TaskListRepository
	logger *slog.Logger
}

func New(r repository.TaskListRepository, l *slog.Logger) *TaskListService {
	return &TaskListService{
		r:      r,
		logger: l,
	}
}

func (t *TaskListService) GetTaskListByID(ctx context.Context, id int32) (*model.TaskList, error) {
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
		Description: data.Description,
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
		UserID:      data.UserID,
		Title:       data.Title,
		Description: data.Description,
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
	Title       string
	Description *string
	TaskListID  int32
	UserID      int32
	Status      model.TaskStatus
}

func (t *TaskListService) CreateTask(ctx context.Context, data CreateTaskOptions) (int32, error) {
	opts := repository.CreateTaskDTO{
		TaskListID:  data.TaskListID,
		UserID:      data.UserID,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
	}
	id, err := t.r.CreateTask(ctx, opts)

	if err != nil {
		t.logger.Error("CreateTask", "error", err)
		return 0, err
	}

	return id, err
}

func (t *TaskListService) DeleteTask(ctx context.Context, id int32) error {
	err := t.r.DeleteTask(ctx, id)
	if err != nil {
		t.logger.Error("DeleteTask", "error", err)
		return err
	}
	return nil
}

type UpdateTaskOptions struct {
	Title       *string
	Description *string
	TaskListID  *int32
	UserID      *int32
	Status      *model.TaskStatus
}

func (t *TaskListService) UpdateTask(ctx context.Context, taskID int32, data UpdateTaskOptions) error {
	opts := repository.UpdateTaskDTO{
		TaskListID:  data.TaskListID,
		UserID:      data.UserID,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
	}

	err := t.r.UpdateTask(ctx, taskID, opts)
	if err != nil {
		t.logger.Error("UpdateTask", "error", err)
		return err
	}

	return nil
}

type GetTasksOptions struct {
	Limit  int32
	Offset int32
}

func (t *TaskListService) GetTasksByTaskListID(ctx context.Context, taskListID int32, options GetTasksOptions) ([]model.Task, error) {
	opts := repository.PaginationOptions{
		Limit:  options.Limit,
		Offset: options.Offset,
	}

	list, err := t.r.GetTasksByTaskList(ctx, taskListID, opts)

	if err != nil {
		t.logger.Error("GetTasksByTaskListID", "error", err)
		return nil, err
	}

	return list, err
}

func (t *TaskListService) GetTaskByID(ctx context.Context, id int32) (*model.Task, error) {
	task, err := t.r.GetTaskByID(ctx, id)
	if err != nil {
		t.logger.Error("GetTaskByID", "error", err)
		return nil, err
	}
	return task, err
}
