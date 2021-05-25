package verdict

import (
	"encoding/json"
	"time"
)

type Status uint8

const (
	InProgress Status = iota
	Passed
	Failed
	SyntaxError
	CompilationError
	Dropped
)

type Verdict struct {
	solutionId uint64
	Status     Status
	timestamp  int64
}

// New function is a convenient way to construct Verdict object
func New(solutionId uint64) Verdict {
	return Verdict{
		solutionId,
		InProgress,
		time.Now().Unix(),
	}
}

// String method represents Verdict as a string
func (v Verdict) String() (string, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// InProgress method check whether Verdict is InProgress
func (v Verdict) InProgress() bool {
	return v.Status == InProgress && v.timestamp > 0
}

// UpdateTimestamp method helps to update Timestamp in a unified way
func (v *Verdict) UpdateTimestamp() {
	v.timestamp = time.Now().Unix()
}

// SetStatus method sets Status and updates Verdict Timestamp
func (v *Verdict) SetStatus(status Status) {
	v.Status = status
	v.UpdateTimestamp()
}

// GetSolutionId method helps to retrieve solutionId from Verdict
func (v Verdict) GetSolutionId() uint64 {
	return v.solutionId
}
