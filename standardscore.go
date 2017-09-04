package statistics

func StandardScoreSample(values []float64, value float64) float64 {
	return StandardScore(Mean(values), StandardDeviationSample(values), value)
}

func StandardScorePopulation(values []float64, value float64) float64 {
	return StandardScore(Mean(values), StandardDeviationPopulation(values), value)
}

func StandardScore(mean, stdDev, value float64) float64 {
	return (value - mean) / stdDev
}
