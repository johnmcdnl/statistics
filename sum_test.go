package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{values: []float64{96, 100, 106, 114}}, 416},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.args.values)
			assert.InEpsilon(t, tt.want, got, allowedEpsilon)
		})
	}
}
