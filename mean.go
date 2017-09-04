package statistics

func Mean(values []float64) float64 {
	return Sum(values) / float64(len(values))
}
