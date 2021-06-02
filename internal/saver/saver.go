package saver

import (
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/models"
)

type Saver interface {
	Save(solution models.Solution)
	Close()
}

type saver struct {
	capacity            uint
	slice               []models.Solution
	forgetAllOnOverflow bool
	f                   flusher.Flusher
	// TODO: add timer
}

func (s *saver) Save(solution models.Solution) {
	if len(s.slice) == cap(s.slice) {
		if s.forgetAllOnOverflow {
			s.slice = make([]models.Solution, 0, s.capacity)
		} else {
			halfCapacity := s.capacity / 2
			s.slice = s.slice[halfCapacity:]
		}
	}
	s.slice = append(s.slice, solution)
}

func (s *saver) Close() {
	// TODO: flush
}

// TODO: flush on timer timeout

func New(capacity uint, flusher flusher.Flusher, forgetAllOnOverflow bool) Saver {
	return &saver{
		capacity:            capacity,
		slice:               make([]models.Solution, 0, capacity),
		forgetAllOnOverflow: forgetAllOnOverflow,
		f:                   flusher,
	}
}
