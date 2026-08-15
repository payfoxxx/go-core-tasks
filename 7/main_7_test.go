package main

import "testing"

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	result := fanIn(ch1, ch2)

	expected := map[int]bool {1: true, 2: true, 3: true, 4: true}

	for res := range result {
		if !expected[res] {
			t.Errorf("Unexpected value: %d, want = %v", res, expected)
		} else {
			delete(expected, res)
		}
	}

	if len(expected) > 0 {
		t.Errorf("Not found %v", expected)
	}
}