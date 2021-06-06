// Package utils contains ad hoc functions not included into any other package
package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"reflect"
)

// SplitSolutionsToBatches function splits slice of Solution into batches of predefined number of elements
func SplitSolutionsToBatches(solutions []models.Solution, batchSize int) ([][]models.Solution, error) {
	if solutions == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize = %v doesn't make sense", batchSize)
	}

	batchesCap := len(solutions) / batchSize
	if len(solutions)%batchSize != 0 {
		batchesCap += 1
	}
	batches := make([][]models.Solution, 0, batchesCap)

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
func SplitVerdictsToBatches(verdicts []models.Verdict, batchSize int) ([][]models.Verdict, error) {
	if verdicts == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize = %v doesn't make sense", batchSize)
	}

	batchesCap := len(verdicts) / batchSize
	if len(verdicts)%batchSize != 0 {
		batchesCap += 1
	}
	batches := make([][]models.Verdict, 0, batchesCap)

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
func InvertSolutionsMap(original map[uint64]models.Solution) (map[models.Solution]uint64, error) {
	if original == nil {
		return nil, nil
	}

	inverted := make(map[models.Solution]uint64, len(original))
	for key, value := range original {
		if _, found := inverted[value]; found {
			return nil, fmt.Errorf("can't invert original map, got duplicated value: \"%v\"", value)
		}
		inverted[value] = key
	}
	return inverted, nil
}

// ConvertSolutionsSliceToMap function converts slice of Solution to map Solution.id->Solution
func ConvertSolutionsSliceToMap(original []models.Solution) map[uint64]models.Solution {
	if original == nil {
		return nil
	}

	converted := make(map[uint64]models.Solution, len(original))
	for _, value := range original {
		converted[value.Id()] = value
	}
	return converted
}

// ConvertVerdictsSliceToMap function converts slice of Verdict to map Verdict.SolutionId->Verdict
func ConvertVerdictsSliceToMap(original []models.Verdict) map[uint64]models.Verdict {
	if original == nil {
		return nil
	}

	converted := make(map[uint64]models.Verdict, len(original))
	for _, value := range original {
		converted[value.SolutionId()] = value
	}
	return converted
}

// FilterSolutions function filters in elements from original slice not found in filterOut elements
func FilterSolutions(original []models.Solution, filterOut map[models.Solution]struct{}) []models.Solution {
	if original == nil {
		return nil
	}

	if filterOut == nil {
		return original
	}

	filtered := make([]models.Solution, 0, len(original))
	for _, value := range original {
		if _, found := filterOut[value]; !found {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// ApplySolutionsFilters function applies all necessary filters to original slice
func ApplySolutionsFilters(original []models.Solution) []models.Solution {
	// excluded is a hardcoded set of elements to filter out from original slice
	excluded := map[models.Solution]struct{}{
		{}: {},
	}
	filtered := FilterSolutions(original, excluded)
	return filtered
}

func IsNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}
