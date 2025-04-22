package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/model"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/repository"
	sqlcproj "github.com/yazmeyaa/future_undefined_app/projects/internal/repository/postgres/sqlc/generated"
)

type pgRepository struct {
	q *sqlcproj.Queries
}

func toPgText(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  len(s) > 0,
	}
}

func (p *pgRepository) Create(ctx context.Context, project *model.Project) (int64, error) {
	id, err := p.q.CreateProject(ctx, sqlcproj.CreateProjectParams{
		Title:       project.Title,
		Description: project.Description,
		GithubUrl:   toPgText(project.GithubUrl),
		ProjectUrl:  toPgText(project.ProjectUrl),
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *pgRepository) List(ctx context.Context, limit int32, offset int32) ([]model.Project, error) {
	var l int32 = 10
	var o int32 = 0
	if limit > 0 {
		l = limit
	}
	if offset > 0 {
		o = offset
	}
	list, err := p.q.ListProjects(ctx, sqlcproj.ListProjectsParams{
		Limit:  l,
		Offset: o,
	})
	if err != nil {
		return nil, err
	}

	res := make([]model.Project, len(list))
	if len(list) == 0 {
		return res, nil
	}

	for idx, item := range list {
		res[idx] = model.Project{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			GithubUrl:   item.GithubUrl.String,
			ProjectUrl:  item.ProjectUrl.String,
			CreatedAt:   item.CreatedAt.Time,
		}
	}

	return res, nil
}

func (p *pgRepository) GetById(ctx context.Context, id int64) (*model.Project, error) {
	proj, err := p.q.GetProject(ctx, id)
	if err != nil {
		return nil, err
	}

	res := model.Project{
		ID:          proj.ID,
		Title:       proj.Title,
		Description: proj.Description,
		GithubUrl:   proj.GithubUrl.String,
		ProjectUrl:  proj.ProjectUrl.String,
		CreatedAt:   proj.CreatedAt.Time,
	}

	return &res, nil
}

func (p *pgRepository) Delete(ctx context.Context, id int64) error {
	return p.q.DeleteProject(ctx, id)
}

func (p *pgRepository) Update(ctx context.Context, project *model.Project) error {
	updates := sqlcproj.UpdateProjectParams{
		Title:       project.Title,
		Description: project.Description,
		GithubUrl:   toPgText(project.GithubUrl),
		ProjectUrl:  toPgText(project.ProjectUrl),
		ID:          project.ID,
	}

	return p.q.UpdateProject(ctx, updates)
}

func NewRepository(q *sqlcproj.Queries) repository.ProjectsRepository {
	return &pgRepository{
		q: q,
	}
}
