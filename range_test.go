package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterQuartileRange(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"unsorted", args{values: []float64{11, 7, 1, 3, 4, 5, 5, 6}}, 3},
		{"even", args{values: []float64{3, 5, 7, 8, 9, 11, 15, 16, 20, 21}}, 9},
		{"odd", args{values: []float64{27, 1, 2, 5, 6, 7, 9, 12, 15, 18, 19}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InterQuartileRange(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon, fmtError(tt.args.values, tt.want, got))
		})
	}
}
