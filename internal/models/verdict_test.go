package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVerdict_String(t *testing.T) {
	{
		ptr := NewVerdict(0, 0, InProgress, "")
		verdictStr, err := ptr.String()
		assert.Nil(t, err)
		// check json string without timestamp
		{
			lhs := `{"solution_id":0,"user_id":0,"status":0`
			rhs := `,"comment":""}`
			assert.Contains(t, verdictStr, lhs)
			assert.Contains(t, verdictStr, rhs)
		}
		unmarshalled := &Verdict{}
		err = json.Unmarshal([]byte(verdictStr), unmarshalled)
		assert.Nil(t, err)
		assert.Equal(t, uint64(0), unmarshalled.solutionId)
		assert.Equal(t, uint64(0), unmarshalled.userId)
		assert.Equal(t, InProgress, unmarshalled.status)
		assert.Equal(t, "", unmarshalled.comment)
	}

	{
		ptr := NewVerdict(1, 2, Passed, "Great job!")
		verdictStr, err := ptr.String()
		assert.Nil(t, err)
		// check json string without timestamp
		{
			lhs := `{"solution_id":1,"user_id":2,"status":1`
			rhs := `,"comment":"Great job!"}`
			assert.Contains(t, verdictStr, lhs)
			assert.Contains(t, verdictStr, rhs)
		}
		unmarshalled := &Verdict{}
		err = json.Unmarshal([]byte(verdictStr), unmarshalled)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), unmarshalled.solutionId)
		assert.Equal(t, uint64(2), unmarshalled.userId)
		assert.Equal(t, Passed, unmarshalled.status)
		assert.Equal(t, "Great job!", unmarshalled.comment)
	}
}

func TestVerdict_InProgress(t *testing.T) {
	// in progress
	{
		ptr := NewVerdict(0, 0, InProgress, "")
		inProgress := ptr.InProgress()
		assert.Equal(t, true, inProgress)
	}

	// not in progress
	{
		ptr := NewVerdict(0, 0, Passed, "")
		inProgress := ptr.InProgress()
		assert.Equal(t, false, inProgress)
	}
}

func TestVerdict_UpdateTimestamp(t *testing.T) {
	// nil verdict ptr
	{
		var ptr *Verdict = nil
		assert.NotPanics(t, func() {
			ptr.UpdateTimestamp()
		})
	}

	// non-nil verdict ptr
	{
		ptr := NewVerdict(0, 0, Passed, "")
		start := ptr.timestamp
		time.Sleep(1 * time.Nanosecond)
		ptr.UpdateTimestamp()
		end := ptr.timestamp
		assert.Greater(t, end, start)
	}
}

func TestVerdict_UpdateStatus(t *testing.T) {
	// nil verdict ptr
	{
		var ptr *Verdict = nil
		assert.NotPanics(t, func() {
			ptr.UpdateStatus(InProgress, "", 0)
		})
	}

	// non-nil verdict ptr
	{
		ptr := NewVerdict(1, 2, Passed, "")
		assert.Equal(t, ptr.solutionId, uint64(1))
		assert.Equal(t, ptr.userId, uint64(2))
		assert.Equal(t, ptr.status, Passed)
		assert.Equal(t, ptr.comment, "")
		start := ptr.timestamp
		time.Sleep(1 * time.Nanosecond)
		ptr.UpdateStatus(Failed, "Comment string", 3)
		assert.Equal(t, ptr.solutionId, uint64(1))
		assert.Equal(t, ptr.userId, uint64(3))
		assert.Equal(t, ptr.status, Failed)
		assert.Equal(t, ptr.comment, "Comment string")
		end := ptr.timestamp
		assert.Greater(t, end, start)
	}
}

func TestVerdict_Status(t *testing.T) {
	ptr := NewVerdict(1, 2, Passed, "Great job!")
	status, comment, userId := ptr.Status()
	assert.Equal(t, status, Passed)
	assert.Equal(t, comment, "Great job!")
	assert.Equal(t, userId, uint64(2))
}

func TestVerdict_SolutionId(t *testing.T) {
	ptr := NewVerdict(1, 2, Passed, "Great job!")
	solutionId := ptr.SolutionId()
	assert.Equal(t, solutionId, uint64(1))
}
