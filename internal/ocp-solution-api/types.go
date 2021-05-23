// TODO add ocp-solution-api package description
package ocp_solution_api

type LanguageType uint8

const (
	Go LanguageType = iota
	Cpp
	Python
)

type Solution struct {
	Id        uint64
	Source    string
	Language  LanguageType
	Timestamp int64
}

type VerdictStatus uint8

const (
	Passed VerdictStatus = iota
	Failed
	SyntaxError
	CompilationError
	InProgress
)

type Verdict struct {
	SolutionId uint64
	Status     VerdictStatus
	Timestamp  int64
}
