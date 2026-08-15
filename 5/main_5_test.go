package main

import (
	"testing"
	"slices"
)

func TestIntersectionSlices(t *testing.T) {
	tests := []struct{
		input1 		[]int
		input2 		[]int
		wantSlice	[]int
		wantBool	bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 5, 7, 8, 9},
			[]int{2, 5},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5},
			[]int{},
			false,
		},
	}

	for _, tt := range tests {
		ok, result := intersectionSlices(tt.input1, tt.input2)
		if ok != tt.wantBool || !slices.Equal(result, tt.wantSlice) {
			t.Errorf("gotBool = %v, wantBool = %v; gotResult = %v, wantResult = %v", ok, tt.wantBool, result, tt.wantSlice)
		}
	}

}