package main

import (
	"testing"
)

func TestGenerateString(t *testing.T) {
	secretKey := "testSecretKey"
	expected := "6IytPBf"

	result := generateString(secretKey, "1", 7)

	if len(result) != 7 {
		t.Errorf("Expected string of length 7, but got %d", len(result))
	}

	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestGenerateStringWithDifferentInputs(t *testing.T) {
	secretKey := "testSecretKey"
	result1 := generateString(secretKey, "1", 7)
	result2 := generateString(secretKey, "2", 7)

	if result1 == result2 {
		t.Errorf("Expected different strings for different inputs, but got '%s' and '%s'", result1, result2)
	}
}

func TestGenerateStringWithDifferentLengths(t *testing.T) {
	secretKey := "testSecretKey"

	result1 := generateString(secretKey, "1", 5)
	if len(result1) != 5 {
		t.Errorf("Expected length of 5, but got %d", len(result1))
	}

	result2 := generateString(secretKey, "1", 10)
	if len(result2) != 10 {
		t.Errorf("Expected length of 10, but got %d", len(result2))
	}
}
