package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func dismantlingAction(arr string) byte {
	bytes := []byte(arr)
	length := len(bytes)
	dict := make(map[byte]int64, length)
	for _, b := range bytes {
		dict[b]++
	}
	for _, letter := range bytes {
		if dict[letter] == 1 {
			return letter
		}
	}
	return ' '
}

func TestDismantlingAction(t *testing.T) {
	testCases := []struct {
		name      string
		arr       string
		expectRes byte
	}{
		{
			name:      "第一个元素只出现一次",
			arr:       "abbccdeff",
			expectRes: 'a',
		},
		{
			name:      "不存在",
			arr:       "ccdd",
			expectRes: ' ',
		},
		{
			name:      "一般情况",
			arr:       "loveleetcode",
			expectRes: 'v',
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := dismantlingAction(tc.arr)
			assert.Equal(t, tc.expectRes, res)
		})
	}
}
