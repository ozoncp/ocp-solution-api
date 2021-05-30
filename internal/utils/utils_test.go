package utils

import (
	"github.com/ozoncp/ocp-solution-api/internal/solution"
	"github.com/ozoncp/ocp-solution-api/internal/verdict"
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
		batches, err := SplitSolutionsToBatches([]Solution{}, 0)
		assert.Nil(t, batches)
		assert.NotNil(t, err)
	}

	// fully filled batches
	{
		solutions := []Solution{
			*solution.New(1, 1),
			*solution.New(2, 1),
			*solution.New(3, 1),
			*solution.New(4, 1),
		}
		expected := [][]Solution{
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
		solutions := []Solution{
			*solution.New(1, 1),
			*solution.New(2, 1),
			*solution.New(3, 1),
			*solution.New(4, 1),
			*solution.New(5, 1),
		}
		expected := [][]Solution{
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
		solutions := []Solution{}
		expected := [][]Solution{}
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
		batches, err := SplitVerdictsToBatches([]Verdict{}, 0)
		assert.Nil(t, batches)
		assert.NotNil(t, err)
	}

	// fully filled batches
	{
		verdicts := []Verdict{
			*verdict.New(1, 1, verdict.InProgress, ""),
			*verdict.New(2, 1, verdict.InProgress, ""),
			*verdict.New(3, 1, verdict.InProgress, ""),
			*verdict.New(4, 1, verdict.InProgress, ""),
		}
		expected := [][]Verdict{
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
		verdicts := []Verdict{
			*verdict.New(1, 1, verdict.InProgress, ""),
			*verdict.New(2, 1, verdict.InProgress, ""),
			*verdict.New(3, 1, verdict.InProgress, ""),
			*verdict.New(4, 1, verdict.InProgress, ""),
			*verdict.New(5, 1, verdict.InProgress, ""),
		}
		expected := [][]Verdict{
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
		verdicts := []Verdict{}
		expected := [][]Verdict{}
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
		solutions := map[uint64]Solution{}
		expected := map[Solution]uint64{}
		inverted, err := InvertSolutionsMap(solutions)
		assert.NotNil(t, inverted)
		assert.Nil(t, err)
		assert.Equal(t, expected, inverted)
	}

	// non-empty solutions map
	{
		solutions := map[uint64]Solution{
			1: *solution.New(1, 1),
			2: *solution.New(2, 1),
			3: *solution.New(3, 1),
		}
		expected := map[Solution]uint64{
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
		solutions := map[uint64]Solution{
			1: *solution.New(1, 1),
			2: *solution.New(2, 1),
			3: *solution.New(2, 1),
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
		solutions := []Solution{}
		expected := map[uint64]Solution{}
		converted := ConvertSolutionsSliceToMap(solutions)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}

	// non-empty solutions
	{
		solutions := []Solution{
			*solution.New(1, 1),
			*solution.New(2, 1),
			*solution.New(3, 1),
		}
		expected := map[uint64]Solution{
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
		verdicts := []Verdict{}
		expected := map[uint64]Verdict{}
		converted := ConvertVerdictsSliceToMap(verdicts)
		assert.NotNil(t, converted)
		assert.Equal(t, expected, converted)
	}

	// non-empty solutions
	{
		verdicts := []Verdict{
			*verdict.New(1, 1, verdict.InProgress, ""),
			*verdict.New(2, 1, verdict.InProgress, ""),
			*verdict.New(3, 1, verdict.InProgress, ""),
		}
		expected := map[uint64]Verdict{
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
		solutions := []Solution{}
		filtered := FilterSolutions(solutions, nil)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// empty filter
	{
		solutions := []Solution{}
		filter := map[Solution]struct{}{}
		filtered := FilterSolutions(solutions, filter)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// non-empty filter
	{
		solutions := []Solution{
			*solution.New(1, 1),
			*solution.New(2, 1),
			*solution.New(3, 1),
		}
		filter := map[Solution]struct{}{
			solutions[1]: {},
		}
		expected := []Solution{
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
		solutions := []Solution{}
		filtered := ApplySolutionsFilters(solutions)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}

	// non-empty solutions
	{
		solutions := []Solution{
			*solution.New(1, 1),
			*solution.New(2, 1),
			*solution.New(3, 1),
		}
		filtered := ApplySolutionsFilters(solutions)
		assert.NotNil(t, filtered)
		assert.Equal(t, solutions, filtered)
	}
}
