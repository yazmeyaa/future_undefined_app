package projectsGRPCHandler

import (
	"context"
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProjectHandler struct {
	grpcProjects.UnimplementedProjectsServiceServer
	projectService *service.ProjectsService
}

func (p ProjectHandler) Create(ctx context.Context, request *grpcProjects.CreateProjectRequest) (*grpcProjects.CreateProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectHandler) GetById(ctx context.Context, request *grpcProjects.GetProjectByIdRequest) (*grpcProjects.GetProjectByIdResponse, error) {
	proj, err := p.projectService.GetById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	res := grpcProjects.GetProjectByIdResponse{
		Id:          proj.ID,
		Title:       proj.Title,
		Description: proj.Description,
		GithubUrl:   &proj.GithubUrl,
		ProjectUrl:  &proj.ProjectUrl,
		CreatedAt:   timestamppb.New(proj.CreatedAt),
	}

	return &res, err
}

func (p ProjectHandler) GetList(ctx context.Context, request *grpcProjects.GetProjectListRequest) (*grpcProjects.GetProjectListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectHandler) DeleteById(ctx context.Context, request *grpcProjects.DeleteByIdRequest) (*grpcProjects.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectHandler) Update(ctx context.Context, request *grpcProjects.UpdateByIdRequest) (*grpcProjects.UpdateByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectHandler) mustEmbedUnimplementedProjectsServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewGRPCHandler(projectService *service.ProjectsService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}
