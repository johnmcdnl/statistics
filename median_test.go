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
		{"t1", args{values: []float64{96, 100, 106, 114}}, 103},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Median(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon)
		})
	}
}
