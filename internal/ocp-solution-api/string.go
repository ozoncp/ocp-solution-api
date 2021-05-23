package ocp_solution_api

import (
	"encoding/json"
)

func (s Solution) String() string {
	out, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (v Verdict) String() string {
	out, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(out)
}
