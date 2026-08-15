package main

import (
	"testing"
	"slices"
)

func TestExamineSlices(t *testing.T) {
	tests := []struct{
		input1 []string
		input2 []string
		want []string
	}{
		{
			[]string{"a", "b", "c"},
			[]string{"b"},
			[]string{"a", "c"},
		},
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
			[]string{},
		},
	}

	for _, tt := range tests {
		got := examineSlices(tt.input1, tt.input2)
		if !slices.Equal(got, tt.want) {
			t.Errorf("examineSlice = %v, want = %v len(got)=%d, len(want)=%d", got, tt.want, len(got), len(tt.want))
		}
	}
}