package main

import (
	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"hello", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}
	for _, tt := range tests {
		got := getTypeOfVar(tt.input)
		if got != tt.expected {
			t.Errorf("getTypeOfVar(%v) = %s; want = %s", tt.input, got, tt.expected)
		}
	}
}

func TestConvertToStrings(t *testing.T) {
	input := []interface{}{42, 3.14, "Golang", true}
	want := "423.14Golangtrue"
	got := convertToString(input...)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertToString = %v, want = %v", got, want)
	}
}

func TestConvertToRune(t *testing.T) {
	s := "Hello"
	want := []rune{'H', 'e', 'l', 'l', 'o'}
	got := convertToRune(s)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertToRune = %v, want = %v", got, want)
	}
}

func TestHashRunesWithSalt(t *testing.T) {
	rs := []rune("abcd")
	salt := "go-2024"
	expected := "eef2e40f2b77de8474f4dec20ca90bd2e84b9eb3ddb332bd53f57ce69bb784ae"
	hashed := hash(rs, salt)
	if hashed != expected {
		t.Errorf("hash = %v, want = %v", hashed, expected)
	}
}
