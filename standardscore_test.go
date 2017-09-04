package statistics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardScoreSample(t *testing.T) {
	type args struct {
		values []float64
		value  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"standardDeviationSample", args{values: []float64{10, 15, 30, 74, 1, 30}, value: 45.3}, 0.72145},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardScoreSample(tt.args.values, tt.args.value)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}

func TestStandardScorePopulation(t *testing.T) {
	type args struct {
		values []float64
		value  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"standardScorePopulation", args{values: []float64{10, 15, 30, 74, 1, 30}, value: 45.3}, 0.79031},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardScorePopulation(tt.args.values, tt.args.value)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}

func TestStandardScore(t *testing.T) {
	type args struct {
		mean   float64
		stdDev float64
		value  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"standardScore", args{mean: 36.3, stdDev: 9.2, value: 45.3}, 0.97826},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardScore(tt.args.mean, tt.args.stdDev, tt.args.value)
			errMsg := fmt.Sprintf("Mean: %.10g StdDev: %.10g Value: %.10g Want: %.10g Got: %.10g",
				tt.args.mean, tt.args.stdDev, tt.args.value, tt.want, got)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, errMsg)
		})
	}
}
