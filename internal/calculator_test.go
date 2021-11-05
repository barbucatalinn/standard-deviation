//go:build !integration
// +build !integration

package internal

import "testing"

func Test_calculateStandardDeviationV1(t *testing.T) {
	t.Parallel()

	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "default",
			args: args{
				data: []float64{2, 4, 6, 8, 10},
			},
			want: 3.1622776601683795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateStandardDeviationV1(tt.args.data); got != tt.want {
				t.Errorf("calculateStandardDeviationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateStandardDeviationV2(t *testing.T) {
	t.Parallel()

	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "default",
			args: args{
				data: []float64{2, 4, 6, 8, 10},
			},
			want: 3.1622776601683795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateStandardDeviationV2(tt.args.data); got != tt.want {
				t.Errorf("calculateStandardDeviationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
