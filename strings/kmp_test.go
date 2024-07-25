package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseLPS struct {
	word     string
	expected []int
	desc     string
}

func TestBuildLPS(t *testing.T) {
	testCases := []testCaseLPS{
		{"aaa", []int{0, 1, 2}, "all chars the same"},
		{"", []int{}, "empty word"},
		{"AABAACAABAA", []int{0, 1, 0, 1, 2, 0, 1, 2, 3, 4, 5}, "complex #1"},
		{"AAACAAAAAC", []int{0, 1, 2, 0, 1, 2, 3, 3, 3, 4}, "complex #2"},
		{"acacxababacc", []int{0, 0, 1, 2, 0, 1, 0, 1, 0, 1, 2, 0}, "complex #3"},
		{"AAABAAA", []int{0, 1, 2, 0, 1, 2, 3}, "palindrome"},
		{"ABCDE", []int{0, 0, 0, 0, 0}, "all chars different"},
		{"A", []int{0}, "single char"},
		{"abcabcabcabc", []int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, "repeated substrings"},
	}

	for _, tc := range testCases {
		table := buildLPSTable(tc.word)
		assert.Equalf(t, tc.expected, table, "failed on testcase %v", tc.desc)
	}
}

type testCaseKMP struct {
	haystack string
	needle   string
	expected []int
	desc     string
}

func TestKMP(t *testing.T) {
	testCases := []testCaseKMP{
		{"aaaaa", "aa", []int{0, 1, 2, 3}, "basic all matches"},
		{"AAAAABAAABA", "AAAA", []int{0, 1}, "complex #1"},
		{"ABABDABACDABABCABAB", "ABABCABAB", []int{10}, "complex #2"},
		{"AAABAABAABABAA", "ABAA", []int{2, 5, 10}, "complex #3"},
		{"mississippi", "issip", []int{4}, "complex #4"},
	}

	for _, tc := range testCases {
		actual := FindSubstrings(tc.haystack, tc.needle)
		assert.Equalf(t, tc.expected, actual, "failed on testcase %v", tc.desc)
	}
}
