package main

import (
	"testing"
	"time"	
)

func TestRandomGenerator(t *testing.T) {
	ch := randomGenerator()

	select {
	case val := <-ch:
		if val < 0 && val >= 50 {
			t.Errorf("Value not in a specified range")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("block")
	}
}