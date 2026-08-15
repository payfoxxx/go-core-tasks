package main

import (
	"fmt"
	"sync"
	"time"
)

type MyWaitGroup struct {
	mu    sync.Mutex
	count int
	ch    chan struct{}
}

// Add - добавляет счетчик
func (wg *MyWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.count += delta
	if wg.count < 0 {
		panic("negative counter in MyWaitGroup")
	}
	if wg.count == 0 && wg.ch != nil {
		close(wg.ch)
		wg.ch = nil
	}
}

// Done - уменьшает счётчик на 1
func (wg *MyWaitGroup) Done() {
	wg.Add(-1)
}

// Wait - блокировка до обнуления счётчика (ожидание)
func (wg *MyWaitGroup) Wait() {
	wg.mu.Lock()
	if wg.count == 0 {
		wg.mu.Unlock()
		return
	}
	if wg.ch == nil {
		wg.ch = make(chan struct{})
	}
	ch := wg.ch
	wg.mu.Unlock()
	<-ch
}

func main() {
	var wg MyWaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("Все горутины завершены")
}
