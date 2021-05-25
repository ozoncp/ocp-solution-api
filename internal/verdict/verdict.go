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
	status     Status
	timestamp  int64
	comment    string
}

// New function is a convenient way to construct Verdict object
func New(solutionId uint64) Verdict {
	return Verdict{
		solutionId,
		InProgress,
		time.Now().Unix(),
		"",
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
	return v.status == InProgress && v.timestamp > 0
}

// UpdateTimestamp method helps to update timestamp in a unified way
func (v *Verdict) UpdateTimestamp() {
	v.timestamp = time.Now().Unix()
}

// SetStatus method sets status and corresponding comment and updates Verdict timestamp
func (v *Verdict) SetStatus(status Status, comment string) {
	v.status = status
	v.comment = comment
	v.UpdateTimestamp()
}

// GetStatus method helps to retrieve status and corresponding comment from Verdict
func (v Verdict) GetStatus() (Status, string) {
	return v.status, v.comment
}

// GetSolutionId method helps to retrieve solutionId from Verdict
func (v Verdict) GetSolutionId() uint64 {
	return v.solutionId
}
