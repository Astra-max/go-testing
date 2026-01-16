package util

import "testing"

func Test_SkipSixth(t *testing.T) {
	testCases := []struct {
		name     string
		val      string
		expected string
	}{
		{"Test: skip sixth", "123", "Invalid Input"},
		{"Test: skip sixth", "This is the write code", "Thisi thewr tecode"},
		{"Test: skip sixth", "", "Invalid Input"},
		{"Test: skip sixth", "Hello there", "Hello here"},
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
				t.Errorf("\nExpected %v got %v", tt.expected, results)
			}
		})
	}
}
