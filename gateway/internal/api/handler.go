package api

import (
	"encoding/json"
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	grpcTaskList "github.com/yazmeyaa/future_undefined_app/api/proto/task_list"
	api "github.com/yazmeyaa/future_undefined_app/gateway/internal/api/gen"
	"net/http"
	"time"
)

var (
	_ api.ServerInterface = Handler{}
)

type Handler struct {
	projectsClient grpcProjects.ProjectsServiceClient
	tasklistClient grpcTaskList.TaskListServiceClient
}

func (h Handler) CreateTaskList(w http.ResponseWriter, r *http.Request) {
	req := api.CreateTaskListJSONRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// TODO: handle error!
		return
	}
	createRes, err := h.tasklistClient.CreateTaskList(r.Context(), &grpcTaskList.CreateTaskListRequest{
		UserId: int32(req.UserId),
		Name:   req.Title,
	})
	if err != nil {
		// TODO: handle error!
		return
	}

	tl, err := h.tasklistClient.GetTaskListByID(r.Context(), &grpcTaskList.GetTaskListByIDRequest{
		Id: createRes.Id,
	})

	if err != nil {
		// TODO: handle error!
		return
	}

	response := api.TaskList{
		CreatedAt:   tl.TaskList.CreatedAt.AsTime().String(),
		Description: tl.TaskList.Description,
		Id:          int(tl.TaskList.Id),
		Title:       tl.TaskList.Title,
		UserID:      int(tl.TaskList.UserId),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		return
	}

}

func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateTaskByID(w http.ResponseWriter, r *http.Request, taskId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteTaskListByID(w http.ResponseWriter, r *http.Request, taskListId int) {
	_, err := h.tasklistClient.DeleteTask(r.Context(), &grpcTaskList.DeleteTaskRequest{
		Id: int32(taskListId),
	})
	if err != nil {
		// TODO: handle error!
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetTaskListByID(w http.ResponseWriter, r *http.Request, taskListId int) {
	listRes, err := h.tasklistClient.GetTaskListByID(r.Context(), &grpcTaskList.GetTaskListByIDRequest{
		Id: int32(taskListId),
	})

	if err != nil {
		// TODO: handle error!
		return
	}

	res := api.TaskList{
		CreatedAt:   listRes.TaskList.CreatedAt.AsTime().String(),
		Description: listRes.TaskList.Description,
		Id:          int(listRes.TaskList.Id),
		Title:       listRes.TaskList.Title,
		UserID:      int(listRes.TaskList.UserId),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err) // TODO: handle error!
	}
}

func (h Handler) UpdateTaskListByID(w http.ResponseWriter, r *http.Request, taskListId int) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTasksByTaskListID(w http.ResponseWriter, r *http.Request, taskListId int) {
	in := grpcTaskList.GetTaskListRequest{TaskListId: int32(taskListId)}
	res, err := h.tasklistClient.GetListOfTasks(r.Context(), &in)
	if err != nil {
		// TODO: handle error!
		return
	}

	tasks := make([]api.Task, len(res.Tasks))
	for i, t := range res.Tasks {
		tasks[i] = api.Task{
			CreatedAt:   t.CreatedAt.AsTime().String(),
			Description: t.Description,
			Id:          int(t.Id),
			Status:      api.TaskStatus(t.Status.String()),
			Title:       t.Title,
			UserID:      int(t.UserId),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		// TODO: handle error!
		panic(err)
	}
}

func (h Handler) DeleteTaskByID(w http.ResponseWriter, r *http.Request, taskId int) {
	in := grpcTaskList.DeleteTaskRequest{
		Id: int32(taskId),
	}
	_, err := h.tasklistClient.DeleteTask(r.Context(), &in)
	if err != nil {
		// TODO: handle error!
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetTaskByID(w http.ResponseWriter, r *http.Request, taskId int) {
	taskResp, err := h.tasklistClient.GetTaskByID(r.Context(), &grpcTaskList.GetTaskByIDRequest{
		Id: int32(taskId),
	})
	if err != nil {
		// TODO: handle error!
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
		// TODO: handle error!
		panic(err)
	}
}

func (h Handler) GetProjectsList(w http.ResponseWriter, r *http.Request, params api.GetProjectsListParams) {
	grpcInput := grpcProjects.GetProjectListRequest{
		Offset: uint32((*params.Page) * (*params.ItemsPerPage)),
		Limit:  uint32(*params.ItemsPerPage),
	}
	list, err := h.projectsClient.GetList(r.Context(), &grpcInput)
	if err != nil {
		// TODO: handle error!
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		// TODO: handle error!
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
