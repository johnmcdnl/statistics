package statistics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even", args{values: []float64{96, 100, 106, 114}}, 103},
		{"odd", args{values: []float64{3, 5, 12}}, 5},
		{"even unsorted", args{values: []float64{3, 13, 7, 5, 21, 23, 39, 23, 40, 23, 14, 12, 56, 23, 29}}, 23},
		{"odd unsorted", args{values: []float64{3, 13, 7, 5, 21, 23, 23, 40, 23, 14, 12, 56, 23, 29}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Median(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}
