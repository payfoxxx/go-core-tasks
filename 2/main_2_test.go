package main

import (
	"testing"
	"slices"
)

func TestGenerateRandomSlice(t *testing.T) {
	slice1 := generateRandomSlice()
	slice2 := generateRandomSlice()

	if slices.Equal(slice1, slice2) {
		t.Errorf("Slice 1 and Slice 2 are equal")
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input 	 []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5}, []int{}},
	}
	
	for _, tt := range tests {
		got := sliceExample(tt.input)
		if !slices.Equal(tt.expected, got) {
			t.Errorf("sliceExample = %v, want = %v", got, tt.expected)
		}
	}
}

func TestAddElemets(t *testing.T) {
	input := []int {1, 2, 3}
	numberToAdd := 8
	want := []int {1, 2, 3, 8}

	got := addElemets(input, numberToAdd)
	
	if !slices.Equal(got, want) {
		t.Errorf("addElements = %v, want = %v, len(got)= %d, len(want)= %d", got, want, len(got), len(want))
	}
}

func TestCopySlice(t *testing.T) {
	input := []int {1, 2, 3}

	got := copySlice(input)
	if !slices.Equal(got, input) {
		t.Errorf("copySlice = %v, want = %v, len(got)= %d, len(want)= %d", got, input, len(got), len(input))
	}

	input = append(input, 2)
	if slices.Equal(got, input) {
		t.Errorf("Pointer was copied")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int {1, 2, 3}
	removeIndex := 1
	want := []int {1, 3}

	got := removeElement(input, removeIndex)

	if !slices.Equal(want, got) {
		t.Errorf("removeElement = %v, want = %v, len(got)= %d, len(want)= %d", got, want, len(got), len(want))
	}
}