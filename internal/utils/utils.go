// Package utils contains ad hoc functions not included into any other package
package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-solution-api/internal/ocp-solution-api"
)

type Solution ocp_solution_api.Solution
type Verdict ocp_solution_api.Verdict

// SplitSolutionsToBatches function splits slice of Solution into batches of predefined number of elements
// TODO: add test
func SplitSolutionsToBatches(solutions []Solution, batchSize int) [][]Solution {
	if solutions == nil {
		return nil
	}

	if batchSize < 1 {
		panic(fmt.Sprintf("batchSize = %v doesn't make sense", batchSize))
	}

	batchesCap := (len(solutions) / batchSize) + 1
	batches := make([][]Solution, 0, batchesCap)
	for len(solutions) > batchSize {
		batches = append(batches, solutions[:batchSize])
		solutions = solutions[batchSize:]
	}
	if len(solutions) > 0 {
		batches = append(batches, solutions[:])
	}
	return batches
}

// SplitVerdictsToBatches function splits slice of Verdict into batches of predefined number of elements
// TODO: add test
func SplitVerdictsToBatches(verdicts []Verdict, batchSize int) [][]Verdict {
	if verdicts == nil {
		return nil
	}

	if batchSize < 1 {
		panic(fmt.Sprintf("batchSize = %v doesn't make sense", batchSize))
	}

	batchesCap := (len(verdicts) / batchSize) + 1
	batches := make([][]Verdict, 0, batchesCap)
	for len(verdicts) > batchSize {
		batches = append(batches, verdicts[:batchSize])
		verdicts = verdicts[batchSize:]
	}
	if len(verdicts) > 0 {
		batches = append(batches, verdicts[:])
	}
	return batches
}

// InvertSolutionsMap function inverts key->value map to value->key map
// TODO: add test
func InvertSolutionsMap(original map[uint64]Solution) map[Solution]uint64 {
	if original == nil {
		return nil
	}

	inverted := map[Solution]uint64{}
	for key, value := range original {
		if _, found := inverted[value]; found {
			panic(fmt.Sprintf("can't invert original map, got duplicated value: \"%v\"", value))
		}
		inverted[value] = key
	}
	return inverted
}

// ConvertSolutionsSliceToMap function converts slice of Solution to map Solution.id->Solution
// TODO: add test
func ConvertSolutionsSliceToMap(origial []Solution) map[uint64]Solution {
	if origial == nil {
		return nil
	}

	converted := map[uint64]Solution{}
	for _, value := range origial {
		converted[value.Id] = value
	}
	return converted
}

// ConvertVerdictsSliceToMap function converts slice of Verdict to map Verdict.SolutionId->Verdict
// TODO: add test
func ConvertVerdictsSliceToMap(origial []Verdict) map[uint64]Verdict {
	if origial == nil {
		return nil
	}

	converted := map[uint64]Verdict{}
	for _, value := range origial {
		converted[value.SolutionId] = value
	}
	return converted
}

// FilterSolutions function filters in elements from original slice not found in filterOut elements
// TODO: add test
func FilterSolutions(original []Solution, filterOut map[Solution]struct{}) []Solution {
	if original == nil {
		return nil
	}

	if filterOut == nil {
		return original
	}

	filtered := make([]Solution, 0, cap(original))
	for _, value := range original {
		if _, found := filterOut[value]; !found {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

// ApplySolutionsFilters function applies all necessary filters to original slice
// TODO: add test
func ApplySolutionsFilters(original []Solution) []Solution {
	// excluded is a hardcoded set of elements to filter out from original slice
	excluded := map[Solution]struct{}{
		{}: {},
	}
	filtered := FilterSolutions(original, excluded)
	return filtered
}
