package statistics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{values: []float64{96, 100, 106, 114}}, 104},
		{"mean", args{values: []float64{10, 15, 30, 74, 1, 30}}, 26.66667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mean(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon)
		})
	}
}
