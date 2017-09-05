package statistics

func isEven(i int) bool {
	return i%2 == 0
}

func isOdd(i int) bool {
	return !isEven(i)
}

func copySlice(s []float64) []float64 {
	c := make([]float64, len(s))
	copy(c, s)
	return c
}