package utils

import (
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitSolutionsToBatches(t *testing.T) {
	// nil solutions
	{
		batches, err := SplitSolutionsToBatches(nil, 2)
		assert.Nil(t, batches)
		assert.Nil(t, err)
	}

	// invalid batch size
	{
		batches, err := SplitSolutionsToBatches([]models.Solution{}, 0)
		assert.Nil(t, batches)
		assert.NotNil(t, err)
	}

	// fully filled batches
	{
		solutions := []models.Solution{
			*models.NewSolution(1, 1),
			*models.NewSolution(2, 1),
			*models.NewSolution(3, 1),
			*models.NewSolution(4, 1),
		}
		expected := [][]models.Solution{
			{
				solutions[0],
				solutions[1],
			},
			{
				solutions[2],
				solutions[3],
			},
		}
		batches, err := SplitSolutionsToBatches(solutions, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}

	// non-fully filled batches
	{
		solutions := []models.Solution{
			*models.NewSolution(1, 1),
			*models.NewSolution(2, 1),
			*models.NewSolution(3, 1),
			*models.NewSolution(4, 1),
			*models.NewSolution(5, 1),
		}
		expected := [][]models.Solution{
			{
				solutions[0],
				solutions[1],
			},
			{
				solutions[2],
				solutions[3],
			},
			{
				solutions[4],
			},
		}
		batches, err := SplitSolutionsToBatches(solutions, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}

	// empty original slice
	{
		solutions := []models.Solution{}
		expected := [][]models.Solution{}
		batches, err := SplitSolutionsToBatches(solutions, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}
}

func TestSplitVerdictsToBatches(t *testing.T) {
	// nil solutions
	{
		batches, err := SplitVerdictsToBatches(nil, 2)
		assert.Nil(t, batches)
		assert.Nil(t, err)
	}

	// invalid batch size
	{
		batches, err := SplitVerdictsToBatches([]models.Verdict{}, 0)
		assert.Nil(t, batches)
		assert.NotNil(t, err)
	}

	// fully filled batches
	{
		verdicts := []models.Verdict{
			*models.NewVerdict(1, 1, models.InProgress, ""),
			*models.NewVerdict(2, 1, models.InProgress, ""),
			*models.NewVerdict(3, 1, models.InProgress, ""),
			*models.NewVerdict(4, 1, models.InProgress, ""),
		}
		expected := [][]models.Verdict{
			{
				verdicts[0],
				verdicts[1],
			},
			{
				verdicts[2],
				verdicts[3],
			},
		}
		batches, err := SplitVerdictsToBatches(verdicts, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}

	// non-fully filled batches
	{
		verdicts := []models.Verdict{
			*models.NewVerdict(1, 1, models.InProgress, ""),
			*models.NewVerdict(2, 1, models.InProgress, ""),
			*models.NewVerdict(3, 1, models.InProgress, ""),
			*models.NewVerdict(4, 1, models.InProgress, ""),
			*models.NewVerdict(5, 1, models.InProgress, ""),
		}
		expected := [][]models.Verdict{
			{
				verdicts[0],
				verdicts[1],
			},
			{
				verdicts[2],
				verdicts[3],
			},
			{
				verdicts[4],
			},
		}
		batches, err := SplitVerdictsToBatches(verdicts, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}

	// empty original slice
	{
		verdicts := []models.Verdict{}
		expected := [][]models.Verdict{}
		batches, err := SplitVerdictsToBatches(verdicts, 2)
		assert.NotNil(t, batches)
		assert.Nil(t, err)
		assert.Equal(t, expected, batches)
	}
}

func TestInvertSolutionsMap(t *testing.T) {
	// nil solutions map
	{
		inverted, err := InvertSolutionsMap(nil)
		assert.Nil(t, inverted)
		assert.Nil(t, err)
	}

	// empty solutions map
	{
		solutions := map[uint64]models.Solution{}
		expected := map[models.Solution]uint64{}
		inverted, err := InvertSolutionsMap(solutions)
		assert.NotNil(t, inverted)
		assert.Nil(t, err)
		assert.Equal(t, expected, inverted)
	}

	// non-empty solutions map
	{
		solutions := map[uint64]models.Solution{
			1: *models.NewSolution(1, 1),
			2: *models.NewSolution(2, 1),
			3: *models.NewSolution(3, 1),
		}
		expected := map[models.Solution]uint64{
			solutions[1]: 1,
			solutions[2]: 2,
			solutions[3]: 3,
		}
		inverted, err := InvertSolutionsMap(solutions)
		assert.NotNil(t, inverted)
		assert.Nil(t, err)
		assert.Equal(t, expected, inverted)
	}

	// duplicates in solutions map
	{
		solutions := map[uint64]models.Solution{
			1: *models.NewSolution(1, 1),
			2: *models.NewSolution(2, 1),
			3: *models.NewSolution(2, 1),
		}
		inverted, err := InvertSolutionsMap(solutions)
		assert.Nil(t, inverted)
		assert.NotNil(t, err)
	}
}

func TestConvertSolutionsSliceToMap(t *testing.T) {
	// nil solutions
	{
		converted := ConvertSolutionsSliceToMap(nil)
		assert.Nil(t, converted)
	}

	// empty solutions
	{
		solutions := []models.Solution{}
		expected := map[uint64]models.Solution{}
		converted := ConvertSolutionsSliceToMap(solutions)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}

	// non-empty solutions
	{
		solutions := []models.Solution{
			*models.NewSolution(1, 1),
			*models.NewSolution(2, 1),
			*models.NewSolution(3, 1),
		}
		expected := map[uint64]models.Solution{
			1: solutions[0],
			2: solutions[1],
			3: solutions[2],
		}
		converted := ConvertSolutionsSliceToMap(solutions)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}
}

func TestConvertVerdictsSliceToMap(t *testing.T) {
	// nil solutions
	{
		converted := ConvertVerdictsSliceToMap(nil)
		assert.Nil(t, converted)
	}

	// empty solutions
	{
		verdicts := []models.Verdict{}
		expected := map[uint64]models.Verdict{}
		converted := ConvertVerdictsSliceToMap(verdicts)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}

	// non-empty solutions
	{
		verdicts := []models.Verdict{
			*models.NewVerdict(1, 1, models.InProgress, ""),
			*models.NewVerdict(2, 1, models.InProgress, ""),
			*models.NewVerdict(3, 1, models.InProgress, ""),
		}
		expected := map[uint64]models.Verdict{
			1: verdicts[0],
			2: verdicts[1],
			3: verdicts[2],
		}
		converted := ConvertVerdictsSliceToMap(verdicts)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}
}

func TestFilterSolutions(t *testing.T) {
	// nil solutions
	{
		filtered := FilterSolutions(nil, nil)
		assert.Nil(t, filtered)
	}

	// nil filter
	{
		solutions := []models.Solution{}
		filtered := FilterSolutions(solutions, nil)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// empty filter
	{
		solutions := []models.Solution{}
		filter := map[models.Solution]struct{}{}
		filtered := FilterSolutions(solutions, filter)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// non-empty filter
	{
		solutions := []models.Solution{
			*models.NewSolution(1, 1),
			*models.NewSolution(2, 1),
			*models.NewSolution(3, 1),
		}
		filter := map[models.Solution]struct{}{
			solutions[1]: {},
		}
		expected := []models.Solution{
			solutions[0],
			solutions[2],
		}
		filtered := FilterSolutions(solutions, filter)
		assert.NotNil(t, filtered)
		assert.Equal(t, expected, filtered)
	}
}

func TestApplySolutionsFilters(t *testing.T) {
	// nil solutions
	{
		filtered := ApplySolutionsFilters(nil)
		assert.Nil(t, filtered)
	}

	// empty solutions
	{
		solutions := []models.Solution{}
		filtered := ApplySolutionsFilters(solutions)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// non-empty solutions
	{
		solutions := []models.Solution{
			*models.NewSolution(1, 1),
			*models.NewSolution(2, 1),
			*models.NewSolution(3, 1),
		}
		filtered := ApplySolutionsFilters(solutions)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}
}
