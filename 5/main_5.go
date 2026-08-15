package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, res := intersectionSlices(a, b)
	fmt.Printf("Пересечения есть: %v, слайс: %v", ok, res)
}

// intersectionSlices - проверяет есть ли пересечения в слайсах и добавляет их в новый слайс.
func intersectionSlices(slice1, slice2 []int) (bool, []int) {
	result := make([]int, 0)
	hash := make(map[int]bool)

	for _, v := range slice1 {
		hash[v] = true
	}

	for _, v := range slice2 {
		if _, ok := hash[v]; ok {
			result = append(result, v)
		}
	}
	ok := len(result) > 0
	return ok, result
}