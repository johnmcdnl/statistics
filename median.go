package statistics

import "sort"

func Median(values []float64) float64 {

	sort.Float64s(values)
	if isEven(len(values)) {
		return evenMedian(values)
	}
	return oddMedian(values)
}

func evenMedian(values []float64) float64 {
	return Mean(values[len(values)/2-1 : len(values)/2+1])
}

func oddMedian(values []float64) float64 {
	return values[len(values)/2]
}
