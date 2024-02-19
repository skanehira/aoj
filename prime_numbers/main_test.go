package main

import "testing"

func TestPrimer(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{2, 3, 4, 5, 6, 7},
			want:  4,
		},
		{
			input: []int{2, 5, 6, 10, 9, 81, 31},
			want:  3,
		},
		{
			input: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			want:  7,
		},
	}

	for _, tt := range tests {
		got := primer(tt.input)
		if tt.want != got {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}
