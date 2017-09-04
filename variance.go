package statistics

import "math"

func VarianceSample(values []float64) float64 {
	//s2 = Σ ( xi - x )2 / ( n - 1 )

	mean := Mean(values)
	var total float64
	for _, v := range values {
		total += math.Pow(v-mean, 2)
	}
	return total / (float64(len(values)) - 1)

}

func VariancePopulation(values []float64) float64 {
	////σ2 = Σ ( Xi - μ )2 / N
	mean := Mean(values)
	var total float64
	for _, v := range values {
		total += math.Pow(v-mean, 2)
	}
	return total / float64(len(values))
}
