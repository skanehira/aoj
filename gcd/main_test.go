package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		input []int64
		want  int64
	}{
		{
			input: []int64{147, 105},
			want:  21,
		},
		{
			input: []int64{28, 12},
			want:  4,
		},
	}

	for _, tt := range tests {
		got := gcd(tt.input[0], tt.input[1])
		if tt.want != got {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}
