package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarianceSample(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"varianceSample", args{values: []float64{10, 15, 30, 74, 1, 30}}, 667.06667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VarianceSample(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}

func TestVariancePopulation(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"variancePopulation", args{values: []float64{10, 15, 30, 74, 1, 30}}, 555.88889},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VariancePopulation(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}
