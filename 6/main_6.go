package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := randomGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func randomGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(50)
		}
	}()
	return ch
}