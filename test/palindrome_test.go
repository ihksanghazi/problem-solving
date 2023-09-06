package test

import (
	"testing"

	"github.com/ihksanghazi/problem-solving/challenges"
)

func Test_IsPalindrome(t *testing.T) {

	testCase := []struct {
		input    string
		expected bool
	}{
		{"aabbccccbbaa", true},    // Test 1
		{"12ini21", true},         // Test 2
		{"kasur ini rusa", false}, // Test 3
	}

	for _, tc := range testCase {
		actual := challenges.IsPalindrome(tc.input)
		if tc.expected != actual {
			t.Error("'", tc.input, "' is not palindrome")
		}
	}

}
