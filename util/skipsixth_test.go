package util

import "testing"

func Test_SkipSixth() {
	testCases := []struct {
		name     string
		val      string
		expected string
	}{
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "123", "Invalid Input"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			results := SkipSixth(tt.val)
			if results != tt.expected {
				t.Errorf("Expected %v got %v", tt.expected, results)
			}
		})
	}
}
