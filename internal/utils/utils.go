// Package utils contains ad hoc functions not included into any other package
package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-solution-api/internal/solution"
	"github.com/ozoncp/ocp-solution-api/internal/verdict"
)

type Solution = solution.Solution
type Verdict = verdict.Verdict

// SplitSolutionsToBatches function splits slice of Solution into batches of predefined number of elements
func SplitSolutionsToBatches(solutions []Solution, batchSize int) ([][]Solution, error) {
	if solutions == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, errors.New(fmt.Sprintf("batchSize = %v doesn't make sense", batchSize))
	}

	batchesCap := len(solutions) / batchSize
	if len(solutions)%batchSize != 0 {
		batchesCap += 1
	}
	batches := make([][]Solution, 0, batchesCap)

	for len(solutions) > batchSize {
		batches = append(batches, solutions[:batchSize])
		solutions = solutions[batchSize:]
	}
	if len(solutions) > 0 {
		batches = append(batches, solutions[:])
	}
	return batches, nil
}

// SplitVerdictsToBatches function splits slice of Verdict into batches of predefined number of elements
func SplitVerdictsToBatches(verdicts []Verdict, batchSize int) ([][]Verdict, error) {
	if verdicts == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, errors.New(fmt.Sprintf("batchSize = %v doesn't make sense", batchSize))
	}

	batchesCap := len(verdicts) / batchSize
	if len(verdicts)%batchSize != 0 {
		batchesCap += 1
	}
	batches := make([][]Verdict, 0, batchesCap)

	for len(verdicts) > batchSize {
		batches = append(batches, verdicts[:batchSize])
		verdicts = verdicts[batchSize:]
	}
	if len(verdicts) > 0 {
		batches = append(batches, verdicts[:])
	}
	return batches, nil
}

// InvertSolutionsMap function inverts key->value map to value->key map
func InvertSolutionsMap(original map[uint64]Solution) (map[Solution]uint64, error) {
	if original == nil {
		return nil, nil
	}

	inverted := make(map[Solution]uint64, len(original))
	for key, value := range original {
		if _, found := inverted[value]; found {
			return nil, errors.New(fmt.Sprintf("can't invert original map, got duplicated value: \"%v\"", value))
		}
		inverted[value] = key
	}
	return inverted, nil
}

// ConvertSolutionsSliceToMap function converts slice of Solution to map Solution.id->Solution
func ConvertSolutionsSliceToMap(original []Solution) map[uint64]Solution {
	if original == nil {
		return nil
	}

	converted := make(map[uint64]Solution, len(original))
	for _, value := range original {
		converted[value.Id()] = value
	}
	return converted
}

// ConvertVerdictsSliceToMap function converts slice of Verdict to map Verdict.SolutionId->Verdict
func ConvertVerdictsSliceToMap(original []Verdict) map[uint64]Verdict {
	if original == nil {
		return nil
	}

	converted := make(map[uint64]Verdict, len(original))
	for _, value := range original {
		converted[value.SolutionId()] = value
	}
	return converted
}

// FilterSolutions function filters in elements from original slice not found in filterOut elements
func FilterSolutions(original []Solution, filterOut map[Solution]struct{}) []Solution {
	if original == nil {
		return nil
	}

	if filterOut == nil {
		return original
	}

	filtered := make([]Solution, 0, len(original))
	for _, value := range original {
		if _, found := filterOut[value]; !found {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// ApplySolutionsFilters function applies all necessary filters to original slice
func ApplySolutionsFilters(original []Solution) []Solution {
	// excluded is a hardcoded set of elements to filter out from original slice
	excluded := map[Solution]struct{}{
		{}: {},
	}
	filtered := FilterSolutions(original, excluded)
	return filtered
}
