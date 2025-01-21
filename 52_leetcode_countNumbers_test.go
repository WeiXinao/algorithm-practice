package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNumbers(t *testing.T) {
	testCases := []struct {
		name   string
		cnt    int
		expect []int
	}{
		{
			name:   "不大于2位的数字",
			cnt:    2,
			expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := countNumbers(tc.cnt)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func countNumbers(cnt int) []int {
	onePosition := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if cnt == 1 {
		return onePosition
	}
	res := make([]int, len(onePosition))
	lastIndex := 0
	copy(res, onePosition)
	i := 0
	for cnt > 1 {
		newRes := res[lastIndex:]
		for i = 0; i < len(newRes); i++ {
			res = append(res, newRes[i]*10)
			for _, in := range onePosition {
				res = append(res, newRes[i]*10+in)
			}
		}
		lastIndex += i
		cnt--
	}
	return res
}
