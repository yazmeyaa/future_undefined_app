package repository

import (
	"context"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/model"
)

type ProjectsRepository interface {
	Create(ctx context.Context, project *model.Project) (int64, error)
	List(ctx context.Context, limit int32, offset int32) ([]model.Project, error)
	GetById(ctx context.Context, id int64) (*model.Project, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, project *model.Project) error
}
