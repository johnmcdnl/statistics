package statistics

import "testing"

const allowedEpsilon = 0.00001

func Test_isEven(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"zero is even", args{i: 0}, true},
		{"even simple", args{i: 2}, true},
		{"odd simple", args{i: 1}, false},
		{"even larger", args{i: 12698}, true},
		{"odd larger", args{i: 12699}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.i); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOdd(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"zero is even", args{i: 0}, false},
		{"even simple", args{i: 2}, false},
		{"odd simple", args{i: 1}, true},
		{"even larger", args{i: 12698}, false},
		{"odd larger", args{i: 12699}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.i); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
