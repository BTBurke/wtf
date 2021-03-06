package metric

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var _ SeriesRecorder = &SampledSeries{}

type SampledSeries struct {
	s         *Series
	mu        sync.RWMutex
	t         *time.Ticker
	obs       []float64
	transform func([]float64) float64
	done      chan bool
	wg        sync.WaitGroup
	direct    bool
}

func NewSampledSeries(capacity int, sampleWindow time.Duration, transform func([]float64) float64, opts ...SeriesOption) (*SampledSeries, func(), error) {
	s, err := NewSeries(capacity, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create sampled series: %v", err)
	}

	if sampleWindow > 0 {
		ss := &SampledSeries{
			s:         s,
			t:         time.NewTicker(sampleWindow),
			obs:       make([]float64, 0),
			transform: transform,
			done:      make(chan bool),
		}
		ss.wg.Add(1)
		go func(s *SampledSeries) {
			defer s.wg.Done()
			for {
				select {
				case <-s.t.C:
					s.mu.Lock()
					if len(s.obs) == 0 {
						s.s.Record(0.0)
					} else {
						s.s.Record(s.transform(s.obs))
						s.obs = make([]float64, 0)
					}
					s.mu.Unlock()
				case <-s.done:
					s.t.Stop()
					return
				}
			}
		}(ss)
		return ss, func() { ss.done <- true; ss.wg.Wait() }, nil
	} else {
		// if duration is zero, use direct mode where observations are immediately written to the underlying
		// series storage.  The transform in this case is a no op.
		ss := &SampledSeries{
			s:         s,
			t:         nil,
			transform: nil,
			done:      nil,
			direct:    true,
		}
		return ss, func() {}, nil
	}
}

// Reset clears all previous recorded values and the count to zero.  This reuses the same backing slice to reduce
// allocations.  It does not attempt to adjust the timing of the sample window, which may cause the initial value to
// be garbage depending on when reset is called within a sample window. For sufficiently large series, this should not matter.
func (s *SampledSeries) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.obs = make([]float64, 0)
	s.s.Reset()
}

func (s *SampledSeries) Capacity() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s.Capacity()
}

func (s *SampledSeries) Record(obs float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.direct {
		s.s.Record(obs)
	} else {
		s.obs = append(s.obs, obs)
	}
}

func (s *SampledSeries) Values() []float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s.Values()
}

func (s *SampledSeries) Name() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s.Name()
}

func (s *SampledSeries) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s.Count()
}

func SampleAverage(obs []float64) float64 {
	if len(obs) == 0 {
		return 0.0
	}
	return SampleSum(obs) / float64(len(obs))
}

func SampleMin(obs []float64) float64 {
	if len(obs) == 0 {
		return 0.0
	}
	min := obs[0]
	for _, o := range obs {
		min = math.Min(min, o)
	}
	return min
}

func SampleMax(obs []float64) float64 {
	if len(obs) == 0 {
		return 0.0
	}
	max := obs[0]
	for _, o := range obs {
		max = math.Max(max, o)
	}
	return max
}

func SampleSum(obs []float64) float64 {
	sum := 0.0
	for _, o := range obs {
		sum += o
	}
	return sum
}
