package saver

import (
	"fmt"
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"time"
)

type Saver interface {
	Save(solution models.Solution)
	Close() error
}

type saver struct {
	capacity            uint
	slice               []models.Solution
	forgetAllOnOverflow bool
	f                   flusher.Flusher
	doneCh              chan bool
}

func (s *saver) Save(solution models.Solution) {
	if len(s.slice) == cap(s.slice) {
		if s.forgetAllOnOverflow {
			s.slice = make([]models.Solution, 0, s.capacity)
		} else {
			halfCapacity := int(s.capacity / 2)
			s.slice = s.slice[halfCapacity:]
		}
	}
	s.slice = append(s.slice, solution)
}

func (s *saver) Close() error {
	var err error
	if s.slice, err = s.f.Flush(s.slice); err != nil {
		err = fmt.Errorf("lost %v solutions: %w", len(s.slice), err)
	}
	close(s.doneCh)
	return err
}

func New(capacity uint, flusher flusher.Flusher, forgetAllOnOverflow bool) Saver {
	s := &saver{
		capacity:            capacity,
		slice:               make([]models.Solution, 0, capacity),
		forgetAllOnOverflow: forgetAllOnOverflow,
		f:                   flusher,
		doneCh:              make(chan bool),
	}

	const tickerPeriod = 5 * time.Second

	go func() {
		ticker := time.NewTicker(tickerPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.slice, _ = s.f.Flush(s.slice)
			case <-s.doneCh:
				return
			}
		}
	}()

	return s
}
