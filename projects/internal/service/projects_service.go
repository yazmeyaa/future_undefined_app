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

func NewProjectsService(repo repository.ProjectsRepository) *ProjectsService {
	return &ProjectsService{repo: repo}
}
