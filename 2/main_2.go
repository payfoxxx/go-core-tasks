package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := generateRandomSlice()
	fmt.Println("Оригинальныый слайс:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Слайс только с четными числами из оригинального:", evenSlice)

	newSliceWithAddedNum := addElemets(originalSlice, 7)
	fmt.Println("Добавление элемента к слайсу:", newSliceWithAddedNum)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Скопированный слайс:", copiedSlice)

	removedSlice := removeElement(originalSlice, 5)
	fmt.Println("Удаление из слайса", removedSlice)

	fmt.Println("Оригинальный слайс: ", originalSlice)
}

// generateRandomSlice - генерирует слайс из 10 случайных чисел от 0 до 9
func generateRandomSlice() []int {
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(10)
	}
	return slice
}

// sliceExample - принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса
func sliceExample(originalSlice []int)[]int {
	newSlice := make([]int, 0)
	for _, i := range originalSlice {
		if i % 2 == 0 {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

// addElemets - принимает слайс и число и добавляет число в конец слайса. Возвращает новый слайс
func addElemets(originalSlice []int, number int)[]int {
	return append(originalSlice, number)
}

// copySlice - принимает слайс и возвращает его копию.
func copySlice(originalSlice []int)[]int {
	copySlice := make([]int, len(originalSlice))
	copy(copySlice, originalSlice)
	return copySlice
}

// removeElement - удаляет элемент из слайса по индексу. Возвращает новый слайс
func removeElement(slice []int, index int)[]int {
	newSlice := make([]int, 0, len(slice) - 1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

