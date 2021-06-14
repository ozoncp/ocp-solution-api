package api

import (
	"context"
	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-verdict-api"
)

const (
	errVerdictNotFound = "verdict not found"
)

type ocpVerdictApi struct {
	desc.UnimplementedOcpVerdictApiServer
}

func (a *ocpVerdictApi) CreateVerdictV1(
	ctx context.Context,
	req *desc.CreateVerdictV1Request,
) (*desc.CreateVerdictV1Response, error) {
	log.Info().Msg(req.String())
	return &desc.CreateVerdictV1Response{}, nil
}

func (a *ocpVerdictApi) ListVerdictsV1(
	ctx context.Context,
	req *desc.ListVerdictsV1Request,
) (*desc.ListVerdictsV1Response, error) {
	log.Info().Msg(req.String())
	return &desc.ListVerdictsV1Response{}, nil
}

func (a *ocpVerdictApi) UpdateVerdictV1(
	ctx context.Context,
	req *desc.UpdateVerdictV1Request,
) (*desc.UpdateVerdictV1Response, error) {
	log.Info().Msg(req.String())
	return &desc.UpdateVerdictV1Response{}, nil
}

func (a *ocpVerdictApi) RemoveVerdictV1(
	ctx context.Context,
	req *desc.RemoveVerdictV1Request,
) (*desc.RemoveVerdictV1Response, error) {
	log.Info().Msg(req.String())
	return &desc.RemoveVerdictV1Response{}, nil
}

func NewOcpVerdictApi() desc.OcpVerdictApiServer {
	return &ocpVerdictApi{}
}
