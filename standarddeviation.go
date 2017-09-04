package statistics

import "math"

func StandardDeviationSample(values []float64) float64 {
	return math.Sqrt(VarianceSample(values))
}

func StandardDeviationPopulation(values []float64) float64 {
	return math.Sqrt(VariancePopulation(values))
}
