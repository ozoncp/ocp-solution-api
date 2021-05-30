package solution

import (
	"encoding/json"
	"github.com/ozoncp/ocp-solution-api/internal/verdict"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution_Id(t *testing.T) {
	ptr := New(1, 2)
	id := ptr.Id()
	assert.Equal(t, uint64(1), id)
}

func TestSolution_IssueId(t *testing.T) {
	ptr := New(1, 2)
	issueId := ptr.IssueId()
	assert.Equal(t, uint64(2), issueId)
}

func TestSolution_String(t *testing.T) {
	// nil verdict
	{
		ptr := New(0, 0)
		solutionStr, err := ptr.String()
		assert.Nil(t, err)
		expected := `{"id":0,"issue_id":0,"verdict":null}`
		assert.Equal(t, expected, solutionStr)
	}

	// non-nil verdict
	{
		ptr := New(1, 2)
		ptr.verdict = verdict.New(3, 4, verdict.Failed, "Try again!")
		solutionStr, err := ptr.String()
		assert.Nil(t, err)
		// check json string without timestamp
		{
			lhs := `{"id":1,"issue_id":2,"verdict":{"solution_id":3,"user_id":4,"status":2`
			rhs := `,"comment":"Try again!"}}`
			assert.Contains(t, solutionStr, lhs)
			assert.Contains(t, solutionStr, rhs)
		}
		unmarshalled := &Solution{}
		err = json.Unmarshal([]byte(solutionStr), unmarshalled)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), unmarshalled.id)
		assert.Equal(t, uint64(2), unmarshalled.issueId)
		solutionId := unmarshalled.verdict.SolutionId()
		status, comment, userId := unmarshalled.verdict.Status()
		assert.Equal(t, uint64(3), solutionId)
		assert.Equal(t, uint64(4), userId)
		assert.Equal(t, status, verdict.Failed)
		assert.Equal(t, comment, "Try again!")
	}
}

func TestSolution_Verdict(t *testing.T) {
	// nil verdict
	{
		ptr := New(0, 0).Verdict()
		assert.Nil(t, ptr)
	}

	// non-nil verdict
	{
		solution := New(0, 0)
		solution.verdict = verdict.New(1, 2, verdict.InProgress, "")
		ptr := solution.Verdict()
		assert.NotNil(t, ptr)
		assert.Equal(t, ptr.SolutionId(), uint64(1))
		status, comment, userId := ptr.Status()
		assert.Equal(t, status, verdict.InProgress)
		assert.Equal(t, comment, "")
		assert.Equal(t, userId, uint64(2))
	}
}
