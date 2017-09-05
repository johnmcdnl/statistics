package statistics

//SimpleMovingAverage returns the simple moving average,
// which filters out 'noise' from a set of data points
func SimpleMovingAverage(values []float64, periods int) []float64 {

	var movingAverages []float64
	for i := periods; i <= len(values); i++ {
		movingAverages = append(movingAverages, Mean(values[i-periods:i]))
	}

	return movingAverages
}
