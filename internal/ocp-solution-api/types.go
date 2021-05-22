// TODO add ocp-solution-api package description
package ocp_solution_api

type LanguageType uint8

const (
	Go LanguageType = iota
	Cpp
	Python
)

type Solution struct {
	id        uint64
	code      string
	language  LanguageType
	timestamp int64
}

type Status uint8

const (
	Passed Status = iota
	Failed
	SyntaxError
	CompilationError
	InProgress
)

type Verdict struct {
	solutionId uint64
	status     Status
	timestamp  int64
}
