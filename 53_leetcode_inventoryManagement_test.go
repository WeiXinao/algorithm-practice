package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryManagement2(t *testing.T) {
	testCases := []struct {
		name   string
		stock  []int
		cnt    int
		expect []int
	}{
		{
			name:   "原序列无序",
			stock:  []int{2, 5, 7, 4},
			cnt:    1,
			expect: []int{2},
		},
		{
			name:   "原序列有序",
			stock:  []int{0, 2, 3, 6},
			cnt:    2,
			expect: []int{0, 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := inventoryManagement2(tc.stock, tc.cnt)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name   string
		stock  []int
		expect []int
	}{
		{
			name:   "原序列无序",
			stock:  []int{2, 5, 7, 4},
			expect: []int{2, 4, 5, 7},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := mergeSort(tc.stock)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func inventoryManagement2(stock []int, cnt int) []int {
	return mergeSort(stock)[:cnt]
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length == 1 {
		return nums
	}
	mid := length / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	i, j := 0, 0
	res := make([]int, 0, length)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for i < len(left) {
		res = append(res, left[i])
		i++
	}
	for j < len(right) {
		res = append(res, right[j])
		j++
	}
	return res
}
