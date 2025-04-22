package api

import (
	"encoding/json"
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	api "github.com/yazmeyaa/future_undefined_app/gateway/internal/api/gen"
	"net/http"
)

type Handler struct {
	projectsClient grpcProjects.ProjectsServiceClient
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

func NewHandler(projectsClient grpcProjects.ProjectsServiceClient) *Handler {
	return &Handler{
		projectsClient: projectsClient,
	}
}
