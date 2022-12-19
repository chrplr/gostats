package gostats

import (
	"math"
)

// An Histrogram holds counts of values falling within intervals defined by breaks. interval1 =[breaks[0]; breaks[1]( ;interval2=[breaks[1], breaks[2]( ...
type Histogram struct {
	Breaks []float64
	Counts []int
}

func NewHistogram(Breaks []float64) *Histogram {
	h := new(Histogram)
	n := len(Breaks)
	h.Breaks = make([]float64, n)
	copy(h.Breaks, Breaks)
	h.Counts = make([]int, n - 1)
	return h
}

func (h *Histogram) AddValue(x float64) {

	// Just ignore out-of-range values
	n := len(h.Breaks)
	if x < h.Breaks[0] || x > h.Breaks[n - 1] {
		return
	}

	var i int
	for i = 1; i < n; i++ {
		if h.Breaks[i] > x {
			break
		}
	}
	h.Counts[i - 1]++
}

func (h *Histogram) AddValues(vals []float64) {
	for _, v := range vals {
		h.AddValue(v)
	}
}

func (h *Histogram) ResetCounts() {
	for i := 0; i < len(h.Counts); i++ {
		h.Counts[i] = 0
	}
}

func (h *Histogram) GetMinMaxCounts() (min, max int) {
	min = math.MaxInt
	max = -min
	for _, v := range h.Counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}


func GetMidPoints(x []float64) []float64 {
	n := len(x)
	midp := make([]float64, n - 1)
	for i := 0; i < n - 1; i++ {
		midp[i] = (x[i] + x[i + 1]) / 2
	}
	return midp
}


func GetMinMax(vals []float64) (min float64, max float64) {
	min = math.Inf(1)
	max = math.Inf(-1)
	for _, v := range vals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func Linspace(start float64, stop float64, nIntervals int) []float64 {
	vals := make([]float64, nIntervals+1)
	vals[0] = start
	vals[nIntervals] = stop
	delta := (stop - start) / float64(nIntervals)
	for i := 1; i < nIntervals; i++ {
		vals[i] = start + float64(i) * delta
	}
	return vals
}
