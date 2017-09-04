package statistics

import (
	"sort"
)

func InterQuartileRange(values []float64) float64 {
	sort.Float64s(values)
	if isEven(len(values)) {
		return evenInterQuartileRange(values)
	}
	return oddInterQuartileRange(values)
}

func evenInterQuartileRange(values []float64) float64 {
	return Median(values[len(values)/2:]) - Median(values[:len(values)/2])
}

func oddInterQuartileRange(values []float64) float64 {
	return Median(values[len(values)/2+1:]) - Median(values[:len(values)/2])
}
