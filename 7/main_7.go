package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := createChannel()
	ch2 := createChannel()
	ch3 := createChannel()

	results := fanIn(ch1, ch2, ch3)

	for result := range results {
		fmt.Println(result)
	}
}

func createChannel() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

// fanIn - сливает N каналов в один
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range ch {
				out <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}