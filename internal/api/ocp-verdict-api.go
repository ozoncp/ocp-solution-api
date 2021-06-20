package api

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/producer"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-verdict-api"

	opentracing "github.com/opentracing/opentracing-go"
)

type ocpVerdictApi struct {
	desc.UnimplementedOcpVerdictApiServer
	repo      repo.Repo
	batchSize int
	producer  producer.Producer
}

func (a *ocpVerdictApi) MultiCreateVerdictV1(
	ctx context.Context,
	req *desc.MultiCreateVerdictV1Request,
) (*desc.MultiCreateVerdictV1Response, error) {

	multiCreateVerdictV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("MultiCreateVerdictV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	span := opentracing.GlobalTracer().StartSpan("MultiCreateVerdictV1")
	defer span.Finish()

	flusher, err := flusher.New(a.repo, a.batchSize)
	respVerdicts := make([]*desc.Verdict, 0)
	if err != nil {
		for _, solution_id := range req.SolutionIds {
			respVerdicts = append(respVerdicts, &desc.Verdict{SolutionId: solution_id})
		}
		return &desc.MultiCreateVerdictV1Response{Verdicts: respVerdicts}, err
	}

	verdicts := make([]models.Verdict, 0)
	for _, solution_id := range req.SolutionIds {
		verdict := models.NewVerdict(solution_id, 0, 0, "")
		verdicts = append(verdicts, *verdict)
	}
	remaining, err := flusher.FlushVerdicts(opentracing.ContextWithSpan(ctx, span), verdicts)
	for _, verdict := range remaining {
		status, comment, userId, timestamp := verdict.Status()
		respVerdicts = append(
			respVerdicts,
			&desc.Verdict{
				SolutionId: verdict.SolutionId(),
				UserId:     userId,
				Status:     desc.Verdict_Status(status),
				Timestamp:  timestamp,
				Comment:    comment,
			},
		)
	}
	if len(remaining) == 0 && err == nil {
		multiCreateVerdictV1Metrics.Succeeded.Inc()
	}
	return &desc.MultiCreateVerdictV1Response{Verdicts: respVerdicts}, err
}

func (a *ocpVerdictApi) CreateVerdictV1(
	ctx context.Context,
	req *desc.CreateVerdictV1Request,
) (*desc.CreateVerdictV1Response, error) {

	createVerdictV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("CreateVerdictV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	verdict, err := a.repo.AddVerdict(ctx, *models.NewVerdict(req.SolutionId, 0, models.InProgress, ""))
	if verdict != nil && err != nil {
		createVerdictV1Metrics.Succeeded.Inc()
	}

	if verdict == nil {
		verdict = models.NewVerdict(0, 0, 0, "")
	}

	status, comment, userId, timestamp := verdict.Status()

	return &desc.CreateVerdictV1Response{
		Verdict: &desc.Verdict{
			SolutionId: verdict.SolutionId(),
			UserId:     userId,
			Status:     desc.Verdict_Status(status),
			Timestamp:  timestamp,
			Comment:    comment,
		},
	}, err
}

func (a *ocpVerdictApi) ListVerdictsV1(
	ctx context.Context,
	req *desc.ListVerdictsV1Request,
) (*desc.ListVerdictsV1Response, error) {

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))

	verdicts, err := a.repo.ListVerdicts(ctx, req.Limit, req.Offset)
	respVerdicts := make([]*desc.Verdict, 0)
	if err == nil {
		for _, verdict := range verdicts {
			status, comment, userId, timestamp := verdict.Status()

			respVerdict := desc.Verdict{
				SolutionId: verdict.SolutionId(),
				UserId:     userId,
				Status:     desc.Verdict_Status(status),
				Timestamp:  timestamp,
				Comment:    comment,
			}
			respVerdicts = append(respVerdicts, &respVerdict)
		}
	}

	return &desc.ListVerdictsV1Response{
		Verdicts: respVerdicts,
	}, err
}

func (a *ocpVerdictApi) UpdateVerdictV1(
	ctx context.Context,
	req *desc.UpdateVerdictV1Request,
) (*desc.UpdateVerdictV1Response, error) {

	updateVerdictV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("UpdateVerdictV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	verdict := models.NewVerdict(req.SolutionId, req.UserId, models.Status(req.Status), req.Comment)
	err := a.repo.UpdateVerdict(ctx, *verdict)

	if err == nil {
		updateVerdictV1Metrics.Succeeded.Inc()
	}

	return &desc.UpdateVerdictV1Response{Success: err == nil}, err
}

func (a *ocpVerdictApi) RemoveVerdictV1(
	ctx context.Context,
	req *desc.RemoveVerdictV1Request,
) (*desc.RemoveVerdictV1Response, error) {

	removeVerdictV1Metrics.Total.Inc()

	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	if err := a.producer.SendMessage("RemoveVerdictV1", string(jsonStr)); err != nil {
		log.Error().Msg(err.Error())
	}

	err := a.repo.RemoveVerdict(ctx, req.SolutionId)

	if err == nil {
		removeVerdictV1Metrics.Succeeded.Inc()
	}

	return &desc.RemoveVerdictV1Response{Success: err == nil}, err
}

func NewOcpVerdictApi(repo repo.Repo, batchSize int) desc.OcpVerdictApiServer {
	p, err := producer.New()
	if err != nil {
		panic(err.Error())
	}
	return &ocpVerdictApi{repo: repo, batchSize: batchSize, producer: p}
}
