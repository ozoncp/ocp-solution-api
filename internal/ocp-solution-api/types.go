// TODO add ocp-solution-api package description
package ocp_solution_api

type LanguageType uint8

const (
	Unknown LanguageType = iota
	Go
	Cpp
	Python
)

type Solution struct {
	UserId     uint64
	SolutionId uint64
	Language   LanguageType
	Timestamp  int64
	SourceCode string
}

type VerdictStatus uint8

const (
	InProgress VerdictStatus = iota
	Passed
	Failed
	SyntaxError
	CompilationError
	Dropped
)

type Verdict struct {
	SolutionId uint64
	Status     VerdictStatus
	Timestamp  int64
}
