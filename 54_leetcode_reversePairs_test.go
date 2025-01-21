package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePairs(t *testing.T) {
	testCases := []struct {
		name   string
		record []int
		expect int
	}{
		{
			name:   "record = [9, 7, 5, 4, 6]",
			record: []int{9, 7, 5, 4, 6},
			expect: 8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := reversePairs(tc.record)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func reversePairs(record []int) int {
	if len(record) == 0 || len(record) == 1 {
		return 0
	}
	mid := len(record) / 2
	res := reversePairs(record[:mid]) + reversePairs(record[mid:])
	tmp := make([]int, len(record))
	copy(tmp, record)
	i, j, k := 0, mid, 0
	for i < mid && j < len(record) {
		if tmp[i] <= tmp[j] {
			record[k] = tmp[i]
			k++
			i++
		}
		if tmp[i] > tmp[j] {
			record[k] = tmp[j]
			k++
			j++
			res += mid - i
		}
	}
	for i < mid {
		record[k] = tmp[i]
		k++
		i++
	}
	for j < len(record) {
		record[k] = tmp[j]
		k++
		j++
	}
	return res
}
