package flusher

import (
	"context"
	"fmt"

	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	"github.com/ozoncp/ocp-solution-api/internal/utils"
)

type Flusher interface {
	FlushSolutions(ctx context.Context, solutions []models.Solution) ([]models.Solution, error)
	FlushVerdicts(ctx context.Context, solutions []models.Verdict) ([]models.Verdict, error)
}

type flusher struct {
	repo      repo.Repo
	batchSize int
}

// Flush method tries to flush all solutions passed to it and returns remaining solutions and error if error occurred
func (f flusher) FlushSolutions(ctx context.Context, solutions []models.Solution) ([]models.Solution, error) {
	batches, err := utils.SplitSolutionsToBatches(solutions, f.batchSize)
	if err != nil {
		return solutions, err
	}

	for i, batch := range batches {
		if err := f.repo.AddSolutions(ctx, batch); err != nil {
			return solutions[i*f.batchSize:], err
		}
	}

	return nil, nil
}

func (f flusher) FlushVerdicts(ctx context.Context, verdicts []models.Verdict) ([]models.Verdict, error) {
	batches, err := utils.SplitVerdictsToBatches(verdicts, f.batchSize)
	if err != nil {
		return verdicts, err
	}

	for i, batch := range batches {
		if err := f.repo.AddVerdicts(ctx, batch); err != nil {
			return verdicts[i*f.batchSize:], err
		}
	}

	return nil, nil
}

func New(repo repo.Repo, batchSize int) (Flusher, error) {
	if utils.IsNil(repo) {
		return nil, fmt.Errorf("got nil Repo")
	}

	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize < 1 doesn't make sense")
	}

	return &flusher{
		repo:      repo,
		batchSize: batchSize,
	}, nil
}
