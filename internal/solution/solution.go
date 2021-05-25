package solution

import (
	"encoding/json"
	"time"
)

type Language uint8

const (
	Go Language = iota
	Cpp
	Python
)

type Solution struct {
	userId     uint64
	solutionId uint64
	language   Language
	timestamp  int64
	sourceCode string
}

// New function is a convenient way to construct Solution object
func New(userId uint64, solutionId uint64, lang Language, sourceCode string) Solution {
	return Solution{
		userId,
		solutionId,
		lang,
		time.Now().Unix(),
		sourceCode,
	}
}

// String method represents Solution as a string
func (s Solution) String() (string, error) {
	out, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// UpdateTimestamp method helps to update timestamp in a unified way
func (s *Solution) UpdateTimestamp() {
	s.timestamp = time.Now().Unix()
}

// GetSolutionId method helps to retrieve solutionId from Solution
func (s Solution) GetSolutionId() uint64 {
	return s.solutionId
}

// GetUserId method helps to retrieve userId from Solution
func (s Solution) GetUserId() uint64 {
	return s.userId
}

// SetSourceCode method helps to update sourceCode and corresponding language and updates timestamp
func (s *Solution) SetSourceCode(sourceCode string, lang Language) {
	s.sourceCode = sourceCode
	s.language = lang
	s.UpdateTimestamp()
}

// GetSourceCode method helps to retrieve sourceCode and corresponding language
func (s Solution) GetSourceCode() (string, Language) {
	return s.sourceCode, s.language
}
