package models

import (
	"encoding/json"
)

type Solution struct {
	id      uint64
	issueId uint64
}

type jsonSolution struct {
	Id      uint64 `json:"id"`
	IssueId uint64 `json:"issue_id"`
}

// NewSolution function is a convenient way to construct Solution object
func NewSolution(id uint64, issueId uint64) *Solution {
	return &Solution{
		id,
		issueId,
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

func (s Solution) MarshalJSON() ([]byte, error) {
	proxy := &jsonSolution{
		Id:      s.id,
		IssueId: s.issueId,
	}
	return json.Marshal(proxy)
}

func (s *Solution) UnmarshalJSON(b []byte) error {
	proxy := &jsonSolution{}

	if err := json.Unmarshal(b, proxy); err != nil {
		return err
	}

	s.id = proxy.Id
	s.issueId = proxy.IssueId

	return nil
}

// Id method helps to retrieve id from Solution
func (s Solution) Id() uint64 {
	return s.id
}

// IssueId method helps to retrieve issueId from Solution
func (s Solution) IssueId() uint64 {
	return s.issueId
}
