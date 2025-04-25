package main

import (
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	grpcTaskList "github.com/yazmeyaa/future_undefined_app/api/proto/task_list"
	internalApi "github.com/yazmeyaa/future_undefined_app/gateway/internal/api"
	api "github.com/yazmeyaa/future_undefined_app/gateway/internal/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func newProjectsClient() grpcProjects.ProjectsServiceClient {
	conn, err := grpc.NewClient("projects:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	return grpcProjects.NewProjectsServiceClient(conn)
}

func newTaskListClient() grpcTaskList.TaskListServiceClient {
	conn, err := grpc.NewClient("task_list:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	return grpcTaskList.NewTaskListServiceClient(conn)
}

func main() {
	grpcProjClient := newProjectsClient()
	taskListClient := newTaskListClient()
	h := internalApi.NewHandler(grpcProjClient, taskListClient)
	x := api.Handler(h)
	if err := http.ListenAndServe(":8080", x); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
