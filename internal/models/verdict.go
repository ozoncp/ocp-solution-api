package models

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

type jsonVerdict struct {
	SolutionId uint64 `json:"solution_id"`
	UserId     uint64 `json:"user_id"`
	Status     Status `json:"status"`
	Timestamp  int64  `json:"timestamp"`
	Comment    string `json:"comment"`
}

// NewVerdict function is a convenient way to construct Verdict object
func NewVerdict(solutionId uint64, userId uint64, status Status, comment string) *Verdict {
	return &Verdict{
		solutionId,
		userId,
		status,
		time.Now().UnixNano(),
		comment,
	}
}

// String method represents Verdict as a string
func (v Verdict) String() (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (v Verdict) MarshalJSON() ([]byte, error) {
	proxy := &jsonVerdict{
		SolutionId: v.solutionId,
		UserId:     v.userId,
		Status:     v.status,
		Timestamp:  v.timestamp,
		Comment:    v.comment,
	}
	return json.Marshal(proxy)
}

func (v *Verdict) UnmarshalJSON(b []byte) error {
	proxy := &jsonVerdict{}

	if err := json.Unmarshal(b, proxy); err != nil {
		return err
	}

	v.solutionId = proxy.SolutionId
	v.userId = proxy.UserId
	v.status = proxy.Status
	v.timestamp = proxy.Timestamp
	v.comment = proxy.Comment

	return nil
}

// InProgress method check whether Verdict is InProgress
func (v Verdict) InProgress() bool {
	return v.status == InProgress && v.timestamp > 0
}

// UpdateTimestamp method helps to update timestamp in a unified way
func (v *Verdict) UpdateTimestamp() {
	if v == nil {
		return
	}

	v.timestamp = time.Now().UnixNano()
}

// UpdateStatus method sets status, corresponding comment, moderator's id and updates Verdict timestamp
func (v *Verdict) UpdateStatus(status Status, comment string, userId uint64) {
	if v == nil {
		return
	}

	v.status = status
	v.comment = comment
	v.userId = userId
	v.UpdateTimestamp()
}

// Status method helps to retrieve status, corresponding comment and moderator's id from Verdict
func (v Verdict) Status() (Status, string, uint64) {
	return v.status, v.comment, v.userId
}

// SolutionId method helps to retrieve solutionId from Verdict
func (v Verdict) SolutionId() uint64 {
	return v.solutionId
}
