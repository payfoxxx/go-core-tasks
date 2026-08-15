package main

import (
	"fmt"
)

// На вход подаются два неупорядоченных слайса строк.
// Например:
// ```
// slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
// slice2 := []string{"banana", "date", "fig"}
// ```
// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	examinedSlice := examineSlices(slice1, slice2)
	fmt.Println(examinedSlice)
}

// examineSlices - возвращает слайс строк, который содержит элементы первого слайса, но отсутствующие во втором
func examineSlices(slice1 []string, slice2 []string)[]string {
	var result []string
	for _, el1 := range slice1 {
		isExist := false
		for _, el2 := range slice2 {
			if el1 == el2 {
				isExist = true
			}
		}
		if !isExist {
			result = append(result, el1)
		} 
	}
	return result	
}