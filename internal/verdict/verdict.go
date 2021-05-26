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
	userId     uint64 // moderator's UserId
	status     Status
	timestamp  int64
	comment    string
}

// New function is a convenient way to construct Verdict object
func New(solutionId uint64, userId uint64, status Status, comment string) *Verdict {
	return &Verdict{
		solutionId,
		userId,
		status,
		time.Now().Unix(),
		comment,
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

// UpdateStatus method sets status, corresponding comment, moderator's id and updates Verdict timestamp
func (v *Verdict) UpdateStatus(status Status, comment string, userId uint64) {
	v.status = status
	v.comment = comment
	v.userId = userId
	v.UpdateTimestamp()
}

// GetStatus method helps to retrieve status, corresponding comment and moderator's id from Verdict
func (v Verdict) GetStatus() (Status, string, uint64) {
	return v.status, v.comment, v.userId
}

// GetSolutionId method helps to retrieve solutionId from Verdict
func (v Verdict) GetSolutionId() uint64 {
	return v.solutionId
}
