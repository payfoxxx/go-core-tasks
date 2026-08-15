package main

import (
	"fmt"
)

// Сделать конвейер чисел
// Даны два канала.
// В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
// затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.

// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		for i := 1; i <= 5; i++ {
			in <- uint8(i)
		}
		close(in)
	}()

	go pipeline(in, out)
	for v := range out {
		fmt.Println(v)
	}
}

func pipeline(in <-chan uint8, out chan<-float64) {
	for v := range in {
		out <- float64(v) * float64(v) * float64(v)
	}
	close(out)
}