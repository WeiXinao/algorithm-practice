package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
  haystackBytes := []byte(haystack)
	nodeBytes := []byte(needle)
	second := 0
	for i, v := range nodeBytes {
		if second == 0 && v == nodeBytes[0] {
			second = i
		}
	}	
	
	idx := -1
	outer:
	for i := 0; i < len(haystackBytes); i++ {
		if haystackBytes[i] == nodeBytes[0] {
			tmp := i
			j := 1
			for ; j < len(nodeBytes); j++ {
				if i + j >= len(haystackBytes) {
					break outer
				}
				if haystackBytes[i + j] != nodeBytes[j] {
					if second != 0 && j > second {
						i = i + second - 1
					}
					break
				} 
			}
			if j == len(nodeBytes) {
				idx = tmp
				break
			}
		}
	}
	return idx
}

func TestStrStr(t *testing.T) {
	testCases := []struct{
		name string
		haystack string
		needle string
		expect int
	}{
		{
			name: "子串存在",
			haystack: "sadbutsad",
			needle: "sad",
			expect: 0,
		},
		{
			name: "子串存在2",
			haystack: "mississippi",
			needle: "issipi",
			expect: -1,
		},
		{
			name: "子串不存在",
			haystack: "leetcode",
			needle: "leeto",
			expect: -1,
		},
	}
	for _, tc := range testCases {
		res := strStr(tc.haystack, tc.needle)
		assert.Equal(t, tc.expect, res)
	}
}