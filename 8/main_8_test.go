package main

import "testing"

func TestWaitGroup(t *testing.T) {
	var wg MyWaitGroup
	wg.Add(1)

	if wg.count != 1 {
		t.Errorf("counter must be 1 after Add")
	}

	wg.Done()

	if wg.count != 0 {
		t.Error("counter must be 0 after Done")
	}
}
