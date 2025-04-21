package main

import (
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	internalApi "github.com/yazmeyaa/future_undefined_app/gateway/internal/api"
	api "github.com/yazmeyaa/future_undefined_app/gateway/internal/api/gen"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	conn, err := grpc.NewClient("projects:50051")
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	grpcClient := grpcProjects.NewProjectsServiceClient(conn)
	h := internalApi.NewHandler(grpcClient)
	x := api.Handler(h)
	if err := http.ListenAndServe(":8080", x); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
