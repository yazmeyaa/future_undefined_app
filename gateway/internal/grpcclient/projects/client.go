package projects

import (
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client grpcProjects.ProjectsServiceClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:   conn,
		client: grpcProjects.NewProjectsServiceClient(conn),
	}

	return client, nil
}
