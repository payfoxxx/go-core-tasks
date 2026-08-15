package main

import (
	"testing"
	"reflect"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	if !m.Exists("a") || !m.Exists("b") {
		t.Errorf("Add function or exists function is not work")
	}

	if m.Exists("c") {
		t.Errorf("Don't added element in map")
	}

	m.Remove("a")
	if m.Exists("a") {
		t.Errorf("Remove doesn't work")
	}

	copy := m.Copy()
	if !reflect.DeepEqual(copy, m.data) {
		t.Errorf("Copy doesn't work")
	}

	copy["b"] = 5
	if copy["b"] == m.data["b"] {
		t.Errorf("Copy is a pointer of StringIntMap")
	}

	elem, ok := m.Get("b")
	if elem != 2 || !ok {
		t.Errorf("Get doesn't work")
	}
}