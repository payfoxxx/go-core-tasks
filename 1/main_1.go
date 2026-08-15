package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	printTypeOfVar(numDecimal)
	printTypeOfVar(numOctal)
	printTypeOfVar(numHexadecimal)
	printTypeOfVar(pi)
	printTypeOfVar(name)
	printTypeOfVar(isActive)
	printTypeOfVar(complexNum)

	mergedString := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(mergedString)

	runeSlice := convertToRune(mergedString)
	fmt.Println(runeSlice)

	hashedRune := hash(runeSlice, "go-2024")
	fmt.Println(hashedRune)
}

// printTypeOfVar - печатает тип переменной
func printTypeOfVar(v interface{}) {
	fmt.Println(getTypeOfVar(v))
}

// getTypeOfVar - возвращает тип переменной
func getTypeOfVar(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// convertToString - преобразует все переменные в строковый тип и объединяет их в одну строку
func convertToString(v ...interface{}) string {
	var builder strings.Builder
	for _, variable := range v {
		builder.WriteString(fmt.Sprint(variable))
	}
	return builder.String()
}

// convertToRune - преобразует строку в срез рун
func convertToRune(s string) []rune {
	return []rune(s)
}

// hash - хэширует срез рун SHA256, добавив в середину соль
func hash(r []rune, salt string) string {
	hasher := sha256.New()
	buf := make([]byte, utf8.UTFMax)

	mid := len(r) / 2

	for _, rr := range r[:mid] {
		n := utf8.EncodeRune(buf, rr)
		hasher.Write(buf[:n])
	}

	hasher.Write([]byte(salt))

	for _, rr := range r[mid:] {
		n := utf8.EncodeRune(buf, rr)
		hasher.Write(buf[:n])
	}

	return fmt.Sprintf("%x", [32]byte(hasher.Sum(nil)))
}
