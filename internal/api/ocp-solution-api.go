package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"encoding/json"

	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-solution-api"
)

const (
	errSolutionNotFound = "solution not found"
)

type ocpSolutionApi struct {
	desc.UnimplementedOcpSolutionApiServer
}

func (a *ocpSolutionApi) CreateSolutionV1(
	ctx context.Context,
	req *desc.CreateSolutionV1Request,
) (*desc.CreateSolutionV1Response, error) {
	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	return &desc.CreateSolutionV1Response{}, nil
}

func (a *ocpSolutionApi) ListSolutionsV1(
	ctx context.Context,
	req *desc.ListSolutionsV1Request,
) (*desc.ListSolutionsV1Response, error) {
	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	return &desc.ListSolutionsV1Response{}, nil
}

func (a *ocpSolutionApi) DescribeSolutionV1(
	ctx context.Context,
	req *desc.DescribeSolutionV1Request,
) (*desc.DescribeSolutionV1Response, error) {
	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	return &desc.DescribeSolutionV1Response{}, nil
}

func (a *ocpSolutionApi) RemoveSolutionV1(
	ctx context.Context,
	req *desc.RemoveSolutionV1Request,
) (*desc.RemoveSolutionV1Response, error) {
	jsonStr, _ := json.Marshal(req)
	log.Info().Msg(string(jsonStr))
	return &desc.RemoveSolutionV1Response{}, nil
}

func NewOcpSolutionApi() desc.OcpSolutionApiServer {
	return &ocpSolutionApi{}
}
