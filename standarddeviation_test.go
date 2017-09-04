package statistics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStandardDeviationSample(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"standardDeviationSample", args{values: []float64{10, 15, 30, 74, 1, 30}}, 25.82763},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardDeviationSample(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}

func TestStandardDeviationPopulation(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"standardDeviationPopulation", args{values: []float64{10, 15, 30, 74, 1, 30}}, 23.5773},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardDeviationPopulation(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}
