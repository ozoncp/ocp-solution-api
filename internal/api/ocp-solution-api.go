package api

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-solution-api"
)

const (
// errSolutionNotFound = "solution not found"
)

type ocpSolutionApi struct {
	desc.UnimplementedOcpSolutionApiServer
	repo repo.Repo
}

func (a *ocpSolutionApi) CreateSolutionV1(
	ctx context.Context,
	req *desc.CreateSolutionV1Request,
) (*desc.CreateSolutionV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	solution, err := a.repo.AddSolution(ctx, *models.NewSolution(0, req.IssueId))

	return &desc.CreateSolutionV1Response{
		Solution: &desc.Solution{
			SolutionId: solution.Id(),
			IssueId:    solution.IssueId(),
			Verdict:    nil,
		},
	}, err
}

func (a *ocpSolutionApi) ListSolutionsV1(
	ctx context.Context,
	req *desc.ListSolutionsV1Request,
) (*desc.ListSolutionsV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	solutions, err := a.repo.ListSolutions(ctx, req.Limit, req.Offset)
	respSolutions := make([]*desc.Solution, 0)
	if err == nil {
		for _, solution := range solutions {
			respSolution := desc.Solution{
				SolutionId: solution.Id(),
				IssueId:    solution.IssueId(),
				Verdict:    nil,
			}
			respSolutions = append(respSolutions, &respSolution)
		}
	}

	return &desc.ListSolutionsV1Response{
		Solutions: respSolutions,
	}, err
}

func (a *ocpSolutionApi) UpdateSolutionV1(
	ctx context.Context,
	req *desc.UpdateSolutionV1Request,
) (*desc.UpdateSolutionV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	solution := models.NewSolution(req.Solution.SolutionId, req.Solution.IssueId)
	// TODO: update verdict
	err := a.repo.UpdateSolution(ctx, *solution)

	return &desc.UpdateSolutionV1Response{Success: err == nil}, err
}

func (a *ocpSolutionApi) RemoveSolutionV1(
	ctx context.Context,
	req *desc.RemoveSolutionV1Request,
) (*desc.RemoveSolutionV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	err := a.repo.RemoveSolution(ctx, req.SolutionId)

	return &desc.RemoveSolutionV1Response{Success: err == nil}, err
}

func NewOcpSolutionApi(repo repo.Repo) desc.OcpSolutionApiServer {
	return &ocpSolutionApi{repo: repo}
}
