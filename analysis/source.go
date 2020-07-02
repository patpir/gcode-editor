package analysis

import (
	"sort"
)

type Source struct {
	points    map[float64][]float64
	KeyUnit   string
	ValueUnit string
}

func (s *Source) Add(key float64, value float64) {
	if s.points == nil {
		s.points = make(map[float64][]float64)
	}

	s.points[key] = append(s.points[key], value)
}

func (s Source) Keys() []float64 {
	var keys []float64
	for key := range s.points {
		keys = append(keys, key)
	}
	sort.Float64s(keys)
	return keys
}

func (s Source) ValuesAt(key float64) []float64 {
	return s.points[key]
}

func (s Source) FirstAt(key float64) (firstValue float64, ok bool) {
	values := s.points[key]
	if len(values) > 0 {
		return values[0], true
	}
	return 0.0, false
}

func (s Source) LastAt(key float64) (lastValue float64, ok bool) {
	values := s.points[key]
	if len(values) > 0 {
		return values[len(values)-1], true
	}
	return 0.0, false
}
