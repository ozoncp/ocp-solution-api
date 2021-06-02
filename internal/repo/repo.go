package repo

import "github.com/ozoncp/ocp-solution-api/internal/models"

type Repo interface {
	AddSolutions(solutions []models.Solution) error
}
