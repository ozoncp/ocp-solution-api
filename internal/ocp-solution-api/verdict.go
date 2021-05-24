package ocp_solution_api

import (
	"encoding/json"
)

// String method represents Verdict as a string
func (v Verdict) String() (string, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// InProgress method check whether Verdict is in progress
func (v Verdict) InProgress() bool {
	return v.Status == InProgress
}
