package main

import "testing"

func TestCheckID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected bool
	}{
		{"single char repeated", "1111", true},
		{"two chars repeated", "1212", true},
		{"three chars repeated", "123123", true},
		{"no repetition", "1234", false},
		{"partial match", "1213", false},
		{"single character", "1", false}, // single char has no divisor pattern
		{"two different chars", "12", false},
		{"four digit pattern", "12341234", true},
		{"odd length no pattern", "12345", false},
		{"pattern at start only", "121234", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkID(tt.id)
			if result != tt.expected {
				t.Errorf("checkID(%q) = %v; want %v", tt.id, result, tt.expected)
			}
		})
	}
}
