package postgres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository"
	sqlctasklist "github.com/yazmeyaa/future_undefined_app/task_list/internal/repository/postgres/sqlc/generated"
	"log/slog"
)

type Repository struct {
	q   *sqlctasklist.Queries
	log *slog.Logger
}

func (r Repository) CreateTask(ctx context.Context, dto repository.CreateTaskDTO) (int32, error) {
	id, err := r.q.CreateTask(ctx, sqlctasklist.CreateTaskParams{
		UserID:      dto.UserID,
		TaskListID:  dto.TaskListID,
		Title:       dto.Title,
		Description: toPgText(dto.Description),
		Status:      sqlctasklist.Taskstatus(dto.Status),
	})

	return id, err
}

func (r Repository) GetTaskByID(ctx context.Context, id int32) (*model.Task, error) {
	task, err := r.q.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Task{
		ID:          task.ID,
		TaskListID:  task.TaskListID,
		Title:       task.Title,
		Description: task.Description.String,
		UserID:      task.UserID,
		Status:      model.TaskStatus(task.Status),
		CreatedAt:   task.CreatedAt.Time,
	}, nil
}

func (r Repository) UpdateTask(ctx context.Context, id int32, dto repository.UpdateTaskDTO) error {
	return r.q.UpdateTask(ctx, sqlctasklist.UpdateTaskParams{
		Title:       dto.Title,
		Description: pgtype.Text{String: dto.Description.String, Valid: dto.Description.Valid},
		UserID:      dto.UserID,
		TaskListID:  dto.TaskListID,
		Status:      sqlctasklist.Taskstatus(dto.Status),
		ID:          id,
	})
}

func (r Repository) DeleteTask(ctx context.Context, id int32) error {
	return r.q.DeleteTask(ctx, id)
}

func (r Repository) DeleteTaskBatch(ctx context.Context, idList []int32) error {
	return r.q.DeleteTaskBatch(ctx, idList)
}

func (r Repository) GetTaskListByID(ctx context.Context, id int32) (*model.TaskList, error) {
	res, err := r.q.GetTaskListByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.TaskList{
		ID:          res.ID,
		UserID:      res.UserID,
		Title:       res.Title,
		Description: res.Description.String,
		CreatedAt:   res.CreatedAt.Time,
	}, nil
}

func (r Repository) GetTasksByTaskList(ctx context.Context, taskListID int32, options repository.PaginationOptions) ([]model.Task, error) {
	res, err := r.q.GetTasksByTaskListID(ctx, sqlctasklist.GetTasksByTaskListIDParams{
		TaskListID: taskListID,
		Limit:      options.Limit,
		Offset:     options.Offset,
	})
	if err != nil {
		return nil, err
	}

	tasks := make([]model.Task, len(res))
	for i, item := range res {
		tasks[i] = model.Task{
			ID:          item.ID,
			TaskListID:  item.TaskListID,
			Title:       item.Title,
			Description: item.Description.String,
			UserID:      item.UserID,
			Status:      model.TaskStatus(item.Status),
			CreatedAt:   item.CreatedAt.Time,
		}
	}

	return tasks, nil
}

func (r Repository) GetTaskListsByUserID(ctx context.Context, userID int32, options repository.PaginationOptions) ([]model.TaskList, error) {
	res, err := r.q.GetTaskListsByUserID(ctx, sqlctasklist.GetTaskListsByUserIDParams{
		UserID: userID,
		Limit:  options.Limit,
		Offset: options.Offset,
	})

	if err != nil {
		return nil, err
	}

	tls := make([]model.TaskList, len(res))
	for i, item := range res {
		tls[i] = model.TaskList{
			ID:          item.ID,
			UserID:      item.UserID,
			Title:       item.Title,
			Description: item.Description.String,
			CreatedAt:   item.CreatedAt.Time,
		}
	}

	return tls, nil
}

func (r Repository) DeleteTaskList(ctx context.Context, id int32) error {
	return r.q.DeleteTaskListByID(ctx, id)
}

func (r Repository) UpdateTaskList(ctx context.Context, id int32, dto repository.UpdateTaskListDTO) error {
	return r.q.UpdateTaskList(ctx, sqlctasklist.UpdateTaskListParams{
		ID:          id,
		UserID:      toPgInt4(dto.UserID),
		Title:       toPgText(dto.Title),
		Description: toPgText(dto.Description),
	})
}

func toPgText(nullString sql.NullString) pgtype.Text {
	return pgtype.Text{
		String: nullString.String,
		Valid:  nullString.Valid,
	}
}

func toPgInt4(nullInt sql.NullInt32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: nullInt.Int32,
		Valid: nullInt.Valid,
	}
}

func NewPostgresqlRepository(q *sqlctasklist.Queries, log *slog.Logger) *Repository {
	return &Repository{
		q:   q,
		log: log,
	}
}
