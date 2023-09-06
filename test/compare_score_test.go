package test

import (
	"testing"

	"github.com/ihksanghazi/problem-solving/challenges"
)

func Test_CompareScore(t *testing.T) {
	testCases := []struct {
		score1   [3]int32
		score2   [3]int32
		expected []int32
	}{
		{[3]int32{1, 2, 3}, [3]int32{3, 2, 1}, []int32{1, 1}},      // Test 1
		{[3]int32{10, 20, 30}, [3]int32{5, 15, 25}, []int32{3, 0}}, // Test 2
		{[3]int32{0, 0, 0}, [3]int32{0, 0, 0}, []int32{0, 0}},      // Test 3
	}

	for _, tc := range testCases {
		actual := challenges.CompareScore(tc.score1, tc.score2)
		if actual[0] != tc.expected[0] && actual[1] != tc.expected[1] {
			t.Errorf("Expected %v, but got %v for score1 %v and score2 %v", tc.expected, actual, tc.score1, tc.score2)
		}
	}
}
