package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_Id(t *testing.T) {
	ptr := NewSolution(1, 2)
	id := ptr.Id()
	assert.Equal(t, uint64(1), id)
}

func TestSolution_IssueId(t *testing.T) {
	ptr := NewSolution(1, 2)
	issueId := ptr.IssueId()
	assert.Equal(t, uint64(2), issueId)
}

func TestSolution_String(t *testing.T) {
	// nil verdict
	{
		ptr := NewSolution(0, 0)
		solutionStr, err := ptr.String()
		assert.Nil(t, err)
		expected := `{"id":0,"issue_id":0}`
		assert.Equal(t, expected, solutionStr)
	}

	// non-nil verdict
	{
		ptr := NewSolution(1, 2)
		solutionStr, err := ptr.String()
		assert.Nil(t, err)
		// check json string without timestamp
		{
			lhs := `{"id":1,"issue_id":2}`
			assert.Contains(t, solutionStr, lhs)
		}
		unmarshalled := &Solution{}
		err = json.Unmarshal([]byte(solutionStr), unmarshalled)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), unmarshalled.id)
		assert.Equal(t, uint64(2), unmarshalled.issueId)
	}
}
