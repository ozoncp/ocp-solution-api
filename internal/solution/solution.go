package solution

import (
	"encoding/json"
	"github.com/ozoncp/ocp-solution-api/internal/verdict"
)

type Verdict = verdict.Verdict

type Solution struct {
	id      uint64
	issueId uint64
	verdict *Verdict
	// TODO: integrate Snippet after 3rd lesson review
	// TODO: integrate Check after 3rd lesson review
}

// New function is a convenient way to construct Solution object
func New(id uint64, issueId uint64) *Solution {
	return &Solution{
		id,
		issueId,
		nil,
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

// Id method helps to retrieve id from Solution
func (s Solution) Id() uint64 {
	return s.id
}

// IssueId method helps to retrieve issueId from Solution
func (s Solution) IssueId() uint64 {
	return s.issueId
}

// Verdict method helps to retrieve verdict from Solution
func (s Solution) Verdict() *Verdict {
	return s.verdict
}
