package service

import (
	"context"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/model"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/repository"
	"log/slog"
)

type ProjectsService struct {
	repo   repository.ProjectsRepository
	logger *slog.Logger
}

func (p *ProjectsService) GetById(ctx context.Context, id int64) (*model.Project, error) {
	proj, err := p.repo.GetById(ctx, id)
	if err != nil {
		p.logger.Warn("ProjectService.GetById", "error", err.Error())
		return nil, err
	}

	return proj, nil
}

type GetProjectsListOptions struct {
	Limit  int32
	Offset int32
}

func (p *ProjectsService) GetList(ctx context.Context, opts GetProjectsListOptions) ([]model.Project, error) {
	list, err := p.repo.List(ctx, opts.Limit, opts.Offset)
	if err != nil {
		p.logger.Warn("ProjectService.GetList", "error", err.Error())
		return nil, err
	}
	return list, nil
}

func NewProjectsService(repo repository.ProjectsRepository, logger *slog.Logger) *ProjectsService {
	return &ProjectsService{repo: repo, logger: logger}
}
