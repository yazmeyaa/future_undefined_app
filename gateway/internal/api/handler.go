package api

import (
	"encoding/json"
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	grpcTaskList "github.com/yazmeyaa/future_undefined_app/api/proto/task_list"
	api "github.com/yazmeyaa/future_undefined_app/gateway/internal/api/gen"
	"net/http"
	"time"
)

type Handler struct {
	projectsClient grpcProjects.ProjectsServiceClient
	tasklistClient grpcTaskList.TaskListServiceClient
}

func (h Handler) DeleteTasksTaskTaskId(w http.ResponseWriter, r *http.Request, taskId int) {
	in := grpcTaskList.DeleteTaskRequest{
		Id: int32(taskId),
	}
	_, err := h.tasklistClient.DeleteTask(r.Context(), &in)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetTasksTaskTaskId(w http.ResponseWriter, r *http.Request, taskId int) {
	taskResp, err := h.tasklistClient.GetTaskByID(r.Context(), &grpcTaskList.GetTaskByIDRequest{
		Id: int32(taskId),
	})
	if err != nil {
		return
	}

	task := taskResp.Task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(api.Task{
		CreatedAt:   task.CreatedAt.AsTime().Format(time.RFC3339),
		Description: task.Description,
		Id:          int(task.Id),
		Status:      api.TaskStatus(task.Status.String()),
		Title:       task.Title,
		UserID:      int(task.UserId),
	}); err != nil {
		panic(err)
	}
}

func (h Handler) PatchTasksTaskTaskId(w http.ResponseWriter, r *http.Request, taskId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteTasksTaskListTaskListId(w http.ResponseWriter, r *http.Request, taskListId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTasksTaskListTaskListId(w http.ResponseWriter, r *http.Request, taskListId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PatchTasksTaskListTaskListId(w http.ResponseWriter, r *http.Request, taskListId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTasksTaskListTaskListIdGetTasks(w http.ResponseWriter, r *http.Request, taskListId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetProjects(w http.ResponseWriter, r *http.Request, params api.GetProjectsParams) {
	grpcInput := grpcProjects.GetProjectListRequest{
		Offset: uint32((*params.Page) * (*params.ItemsPerPage)),
		Limit:  uint32(*params.ItemsPerPage),
	}
	list, err := h.projectsClient.GetList(r.Context(), &grpcInput)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewHandler(projectsClient grpcProjects.ProjectsServiceClient, tasklistClient grpcTaskList.TaskListServiceClient) *Handler {
	return &Handler{
		projectsClient: projectsClient,
		tasklistClient: tasklistClient,
	}
}
