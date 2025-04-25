package grpchandler

import (
	"context"
	grpcTaskList "github.com/yazmeyaa/future_undefined_app/api/proto/task_list"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func modelStatusToGRPCStatus(taskStatus model.TaskStatus) grpcTaskList.TaskStatus {
	switch taskStatus {
	case model.TaskStatusCreated:
		return grpcTaskList.TaskStatus_created
	case model.TaskStatusInProgress:
		return grpcTaskList.TaskStatus_in_progress
	case model.TaskStatusCompleted:
		return grpcTaskList.TaskStatus_completed
	case model.TaskStatusFailed:
		return grpcTaskList.TaskStatus_failed
	case model.TaskStatusCanceled:
		return grpcTaskList.TaskStatus_canceled
	default:
		return grpcTaskList.TaskStatus_created
	}
}

type TaskListHandler struct {
	svc *service.TaskListService
	grpcTaskList.UnimplementedTaskListServiceServer
	logger *slog.Logger
}

func (t TaskListHandler) CreateTask(ctx context.Context, request *grpcTaskList.CreateTaskRequest) (*grpcTaskList.CreateTaskResponse, error) {
	data := service.CreateTaskOptions{
		Title:       request.Title,
		Description: request.Description,
		TaskListID:  request.TaskListId,
		UserID:      request.UserId,
		Status:      model.TaskStatus(request.Status),
	}
	id, err := t.svc.CreateTask(ctx, data)

	if err != nil {
		t.logger.Error("create task error", err)
		return nil, err
	}

	response := &grpcTaskList.CreateTaskResponse{
		Id: id,
	}

	return response, nil

}

func (t TaskListHandler) GetTaskByID(ctx context.Context, request *grpcTaskList.GetTaskByIDRequest) (*grpcTaskList.GetTaskByIDResponse, error) {
	task, err := t.svc.GetTaskByID(ctx, request.Id)

	if err != nil {
		t.logger.Error("get task error", "error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	response := &grpcTaskList.GetTaskByIDResponse{
		Task: &grpcTaskList.Task{
			Id:          task.ID,
			TaskListId:  task.TaskListID,
			Title:       task.Title,
			Description: task.Description,
			UserId:      task.UserID,
			Status:      modelStatusToGRPCStatus(task.Status),
			CreatedAt:   timestamppb.New(task.CreatedAt),
		},
	}

	return response, nil
}

func (t TaskListHandler) GetListOfTasks(ctx context.Context, request *grpcTaskList.GetTaskListRequest) (*grpcTaskList.GetTaskListResponse, error) {
	tasks, err := t.svc.GetTasksByTaskListID(ctx, request.TaskListId, service.GetTasksOptions{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		t.logger.Error("get tasks error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	response := &grpcTaskList.GetTaskListResponse{
		Tasks: make([]*grpcTaskList.Task, 0),
	}

	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		response.Tasks[i] = &grpcTaskList.Task{
			Id:          task.ID,
			TaskListId:  task.TaskListID,
			Title:       task.Title,
			Description: task.Description,
			UserId:      task.UserID,
			CreatedAt:   timestamppb.New(task.CreatedAt),
			Status:      modelStatusToGRPCStatus(task.Status),
		}
	}

	return response, nil
}

func (t TaskListHandler) UpdateTask(ctx context.Context, request *grpcTaskList.UpdateTaskRequest) (*grpcTaskList.UpdateTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskListHandler) DeleteTask(ctx context.Context, request *grpcTaskList.DeleteTaskRequest) (*grpcTaskList.DeleteTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskListHandler) CreateTaskList(ctx context.Context, request *grpcTaskList.CreateTaskListRequest) (*grpcTaskList.CreateTaskListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskListHandler) GetTaskListsByUserID(ctx context.Context, request *grpcTaskList.GetTaskListsByUserIDRequest) (*grpcTaskList.GetTaskListsByUserIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskListHandler) GetTaskListByID(ctx context.Context, request *grpcTaskList.GetTaskListByIDRequest) (*grpcTaskList.GetTaskListByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskListHandler) DeleteTaskList(ctx context.Context, request *grpcTaskList.DeleteTaskListRequest) (*grpcTaskList.DeleteTaskListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewGRPCHandler(taskListService *service.TaskListService, logger *slog.Logger) *TaskListHandler {
	return &TaskListHandler{
		svc:    taskListService,
		logger: logger,
	}
}
