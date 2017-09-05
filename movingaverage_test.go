package statistics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testSlice = []float64{16, 7, 54, 40, 69, 28, 47, 70, 61, 62, 64}
func TestSimpleMovingAverage(t *testing.T) {
	type args struct {
		values  []float64
		periods int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"t1", args{values: copySlice(testSlice), periods: 5}, []float64{37.2, 39.6, 47.6, 50.8, 55, 53.6, 60.8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleMovingAverage(tt.args.values, tt.args.periods)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

