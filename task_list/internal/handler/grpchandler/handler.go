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
		Limit:  request.PageSize,
		Offset: request.PageOffset,
	})
	if err != nil {
		t.logger.Error("get tasks error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	response := &grpcTaskList.GetTaskListResponse{
		Tasks: make([]*grpcTaskList.Task, len(tasks)),
	}

	for i, task := range tasks {
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
	var taskStatus *model.TaskStatus = nil
	if request.Status != nil {
		statusStr := *request.Status
		str := statusStr.String()
		x := model.TaskStatus(str)
		taskStatus = &x
	}
	opts := service.UpdateTaskOptions{
		Title:       request.Title,
		Description: request.Description,
		TaskListID:  request.TaskListId,
		UserID:      request.UserId,
		Status:      taskStatus,
	}

	err := t.svc.UpdateTask(ctx, request.Id, opts)
	if err != nil {
		t.logger.Error("update task error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	task, err := t.svc.GetTaskByID(ctx, request.Id)
	if err != nil {
		t.logger.Error("get task error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &grpcTaskList.UpdateTaskResponse{
		Task: &grpcTaskList.Task{
			Id:          task.ID,
			TaskListId:  task.TaskListID,
			Title:       task.Title,
			Description: task.Description,
			UserId:      task.UserID,
			Status:      modelStatusToGRPCStatus(task.Status),
			CreatedAt:   timestamppb.New(task.CreatedAt),
		},
	}, nil
}

func (t TaskListHandler) DeleteTask(ctx context.Context, request *grpcTaskList.DeleteTaskRequest) (*grpcTaskList.DeleteTaskResponse, error) {
	err := t.svc.DeleteTask(ctx, request.Id)
	if err != nil {
		t.logger.Error("delete task error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &grpcTaskList.DeleteTaskResponse{
		Success: true,
	}, nil
}

func (t TaskListHandler) CreateTaskList(ctx context.Context, request *grpcTaskList.CreateTaskListRequest) (*grpcTaskList.CreateTaskListResponse, error) {
	taskListID, err := t.svc.CreateTaskList(ctx, service.CreateTaskListOptions{
		UserID:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		t.logger.Error("create task list error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	task, err := t.svc.GetTaskListByID(ctx, taskListID)

	if err != nil {
		t.logger.Error("get task list error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	response := &grpcTaskList.CreateTaskListResponse{
		Id:          task.ID,
		UserId:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   timestamppb.New(task.CreatedAt),
	}

	return response, nil
}

func (t TaskListHandler) GetTaskListsByUserID(ctx context.Context, request *grpcTaskList.GetTaskListsByUserIDRequest) (*grpcTaskList.GetTaskListsByUserIDResponse, error) {
	tls, err := t.svc.GetTaskListsByUserID(ctx, request.UserId, service.GetTaskListsOptions{
		Limit:  request.PageSize,
		Offset: request.PageOffset,
	})
	if err != nil {
		t.logger.Error("get task lists by user id error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	response := &grpcTaskList.GetTaskListsByUserIDResponse{
		TaskLists: make([]*grpcTaskList.TaskList, len(tls)),
	}
	for i, tl := range tls {
		response.TaskLists[i] = &grpcTaskList.TaskList{
			Id:          tl.ID,
			UserId:      tl.UserID,
			Title:       tl.Title,
			Description: tl.Description,
			CreatedAt:   timestamppb.New(tl.CreatedAt),
		}
	}

	return response, nil
}

func (t TaskListHandler) GetTaskListByID(ctx context.Context, request *grpcTaskList.GetTaskListByIDRequest) (*grpcTaskList.GetTaskListByIDResponse, error) {
	tl, err := t.svc.GetTaskListByID(ctx, request.Id)
	if err != nil {
		t.logger.Error("get task list by id error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	response := &grpcTaskList.GetTaskListByIDResponse{
		TaskList: &grpcTaskList.TaskList{
			Id:          tl.ID,
			UserId:      tl.UserID,
			Title:       tl.Title,
			Description: tl.Description,
			CreatedAt:   timestamppb.New(tl.CreatedAt),
		},
	}

	return response, nil
}

func (t TaskListHandler) DeleteTaskList(ctx context.Context, request *grpcTaskList.DeleteTaskListRequest) (*grpcTaskList.DeleteTaskListResponse, error) {
	err := t.svc.DeleteTaskList(ctx, request.Id)
	if err != nil {
		t.logger.Error("delete task list error", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &grpcTaskList.DeleteTaskListResponse{
		Success: true,
	}, nil
}

func NewGRPCHandler(taskListService *service.TaskListService, logger *slog.Logger) *TaskListHandler {
	return &TaskListHandler{
		svc:    taskListService,
		logger: logger,
	}
}
