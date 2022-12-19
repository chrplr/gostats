package gostats

import (
	"fmt"
	"math"
	"gonum.org/v1/gonum/stat"
)

type DescriptiveStats struct {
	n                           int
	min, q1, median, q3, max    float64
	mean, sd, se                float64
	absdev, mad, skew, kurtosis float64
}

func (d DescriptiveStats)String() string {
	return fmt.Sprintf("mean = %+v", d.mean)
}

func Desc(data []float64) DescriptiveStats {
	var e DescriptiveStats
	e.n = len(data)
	e.mean = stat.Mean(data, nil)
	e.sd = math.Sqrt(stat.Variance(data, nil))
	//e.median = stat.Quantile(0.5, stat.Empirical, data, nil)
	return e
}
