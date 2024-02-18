package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{4, 6, 1, 3, 8, 10},
			expect: []int{1, 3, 4, 6, 8, 10},
		},
		{
			input:  []int{5, 2, 4, 6, 1, 3},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		got := sort(tt.input)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("unexpected, want: %v, got: %v", tt.expect, got)
		}
	}
}
