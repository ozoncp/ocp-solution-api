package saver

import (
	"fmt"
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/utils"
	"sync"
	"time"
)

type Saver interface {
	Save(solution models.Solution)
	Close() error
}

type saver struct {
	capacity            uint
	mtx                 sync.Mutex
	slice               []models.Solution
	forgetAllOnOverflow bool
	f                   flusher.Flusher
	done                chan struct{}
}

func (s *saver) Save(solution models.Solution) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	var err error
	if s.slice, err = s.f.Flush(s.slice); err != nil {
		err = fmt.Errorf("lost %v solutions: %w", len(s.slice), err)
	}
	s.done <- struct{}{}
	return err
}

func New(capacity uint, flusher flusher.Flusher, forgetAllOnOverflow bool, tickerPeriod time.Duration) (Saver, error) {
	if utils.IsNil(flusher) {
		return nil, fmt.Errorf("got nil Flusher")
	}

	if capacity < uint(1) {
		return nil, fmt.Errorf("zero Saver capacity doesn't make sense")
	}

	s := &saver{
		capacity:            capacity,
		mtx:                 sync.Mutex{},
		slice:               make([]models.Solution, 0, capacity),
		forgetAllOnOverflow: forgetAllOnOverflow,
		f:                   flusher,
		done:                make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(tickerPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.slice, _ = s.f.Flush(s.slice)
			case <-s.done:
				return
			}
		}
	}()

	return s, nil
}
