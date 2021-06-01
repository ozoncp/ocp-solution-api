package repo

import "github.com/ozoncp/ocp-solution-api/internal/solution"

type Solution = solution.Solution

type Repo interface {
	AddSolutions(solution []Solution) error
}
