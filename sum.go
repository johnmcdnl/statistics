package statistics

func Sum(values []float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}
