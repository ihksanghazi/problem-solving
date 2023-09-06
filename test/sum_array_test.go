package test

import (
	"testing"

	"github.com/ihksanghazi/problem-solving/challenges"
)

func Test_SumArray(t *testing.T) {
	testCases := []struct {
		input    []int32
		expected int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 15},       // Test 1
		{[]int32{10, 20, 30, 40, 50}, 150}, // Test 2
		{[]int32{0, 0, 0, 0, 0}, 0},        // Test 3
	}

	for _, tc := range testCases {
		actual := challenges.SumArray(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %d, but got %d for input %v", tc.expected, actual, tc.input)
		}
	}
}
