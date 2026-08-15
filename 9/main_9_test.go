package main

import (
	"testing"
	"slices"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	expected := []float64{1, 8, 27, 64, 125}

	go func() {
		for i := 1; i <= 5; i++ {
			in <- uint8(i)
		}
		close(in)
	}()

	go pipeline(in, out)

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	if !slices.Equal(result, expected) {
		t.Errorf("got=%v, want=%v", result, expected)
	}
}