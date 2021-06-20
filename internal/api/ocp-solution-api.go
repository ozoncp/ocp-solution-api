package api

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/producer"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-solution-api"

	opentracing "github.com/opentracing/opentracing-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ocpSolutionApi struct {
	desc.UnimplementedOcpSolutionApiServer
	repo      repo.Repo
	batchSize int
	producer  producer.Producer
}

func (a *ocpSolutionApi) MultiCreateSolutionV1(
	ctx context.Context,
	req *desc.MultiCreateSolutionV1Request,
) (*desc.MultiCreateSolutionV1Response, error) {

	multiCreateSolutionV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("MultiCreateSolutionV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span := opentracing.GlobalTracer().StartSpan("MultiCreateSolutionV1")
	defer span.Finish()

	flusher, err := flusher.New(a.repo, a.batchSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	solutions := make([]models.Solution, 0)
	for _, issue_id := range req.IssueIds {
		solution := models.NewSolution(0, issue_id)
		solutions = append(solutions, *solution)
	}
	remaining, err := flusher.FlushSolutions(opentracing.ContextWithSpan(ctx, span), solutions)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	multiCreateSolutionV1Metrics.Succeeded.Inc()

	respSolutions := make([]*desc.Solution, 0)
	for _, solution := range remaining {
		respSolutions = append(
			respSolutions,
			&desc.Solution{
				SolutionId: solution.Id(),
				IssueId:    solution.IssueId(),
			},
		)
	}

	return &desc.MultiCreateSolutionV1Response{Solutions: respSolutions}, err
}

func (a *ocpSolutionApi) CreateSolutionV1(
	ctx context.Context,
	req *desc.CreateSolutionV1Request,
) (*desc.CreateSolutionV1Response, error) {

	createSolutionV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("CreateSolutionV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	solution, err := a.repo.AddSolution(ctx, *models.NewSolution(0, req.IssueId))
	if solution == nil || err != nil {
		return nil, status.Error(codes.Aborted, "solution not added: "+err.Error())
	}

	createSolutionV1Metrics.Succeeded.Inc()

	return &desc.CreateSolutionV1Response{
		Solution: &desc.Solution{
			SolutionId: solution.Id(),
			IssueId:    solution.IssueId(),
		},
	}, err
}

func (a *ocpSolutionApi) ListSolutionsV1(
	ctx context.Context,
	req *desc.ListSolutionsV1Request,
) (*desc.ListSolutionsV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	solutions, err := a.repo.ListSolutions(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	respSolutions := make([]*desc.Solution, 0)
	for _, solution := range solutions {
		respSolution := desc.Solution{
			SolutionId: solution.Id(),
			IssueId:    solution.IssueId(),
		}
		respSolutions = append(respSolutions, &respSolution)
	}

	return &desc.ListSolutionsV1Response{
		Solutions: respSolutions,
	}, err
}

func (a *ocpSolutionApi) UpdateSolutionV1(
	ctx context.Context,
	req *desc.UpdateSolutionV1Request,
) (*desc.UpdateSolutionV1Response, error) {

	updateSolutionV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("UpdateSolutionV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	solution := models.NewSolution(req.Solution.SolutionId, req.Solution.IssueId)
	err := a.repo.UpdateSolution(ctx, *solution)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	updateSolutionV1Metrics.Succeeded.Inc()

	return &desc.UpdateSolutionV1Response{Success: true}, err
}

func (a *ocpSolutionApi) RemoveSolutionV1(
	ctx context.Context,
	req *desc.RemoveSolutionV1Request,
) (*desc.RemoveSolutionV1Response, error) {

	removeSolutionV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("RemoveSolutionV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.repo.RemoveSolution(ctx, req.SolutionId)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	removeSolutionV1Metrics.Succeeded.Inc()

	return &desc.RemoveSolutionV1Response{Success: true}, err
}

func NewOcpSolutionApi(repo repo.Repo, batchSize int, producer producer.Producer) desc.OcpSolutionApiServer {
	return &ocpSolutionApi{repo: repo, batchSize: batchSize, producer: producer}
}
