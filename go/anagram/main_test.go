package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		word string
		want []string
	}{
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"ab", "ac", "ba", "bc", "ca", "cb"}},
		{"abcd", []string{"ab", "ac", "ad", "ba", "bc", "bd", "ca", "cb", "cd"}},
	}

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			got := TwoWordedAnagrams(tc.word)

			for _, anagram := range tc.want {
				assert.Contains(t, got, anagram)
			}
		})
	}
}
