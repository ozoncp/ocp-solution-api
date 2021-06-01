package flusher

import (
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	"github.com/ozoncp/ocp-solution-api/internal/solution"
	"github.com/ozoncp/ocp-solution-api/internal/utils"
)

type Solution = solution.Solution

type Flusher interface {
	Flush(solutions []Solution) ([]Solution, error)
}

type flusher struct {
	repo      repo.Repo
	batchSize int
}

// Flush method tries to flush all solutions passed to it and returns remaining solutions and error if error occurred
func (f flusher) Flush(solutions []Solution) ([]Solution, error) {
	batches, err := utils.SplitSolutionsToBatches(solutions, f.batchSize)
	if err != nil {
		return solutions, err
	}

	for i, batch := range batches {
		if err := f.repo.AddSolutions(batch); err != nil {
			return solutions[i*f.batchSize:], err
		}
	}

	return nil, nil
}

func New(repo repo.Repo, batchSize int) Flusher {
	return &flusher{
		repo:      repo,
		batchSize: batchSize,
	}
}
