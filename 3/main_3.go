package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

// Add - добавляет элемент в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove - удаляет элемент по ключу
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy - копирует элементы в новую карту
func (m *StringIntMap) Copy() map[string]int {
	copy := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copy[k] = v
	}
	return copy
}

// Exists - проверяет существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

// Get - получает значение по ключу из карты
func (m *StringIntMap) Get(key string) (int, bool) {
	elem, ok := m.data[key]
	return elem, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("Hello", 2)
	m.Add("Privet", 3)
	fmt.Println("Исходная карта", m.data)

	m.Remove("Hello")
	fmt.Println("После удаления", m.data)

	newMap := m.Copy()
	fmt.Println("Скопированная карта", newMap)

	fmt.Println("Элемент Privet существует", m.Exists("Privet"))

	elem, ok := m.Get("Privet")
	fmt.Printf("Элемент %s, сущесвует: %v, значение: %d", "Privet", ok, elem)
}