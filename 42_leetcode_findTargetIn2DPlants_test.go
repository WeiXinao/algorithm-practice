package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findTargetIn2DPlants(plants [][]int, target int) bool {
	if len(plants) == 0 {
		return false
	}
	i, j := 0, len(plants[0])-1
	for i <= j {
		mid := (i + j) / 2
		if target < plants[0][mid] {
			j = mid - 1
		} else if target > plants[0][mid] {
			i = mid + 1
		} else {
			return true
		}
	}

	columnEnd := j

	for k := 0; k <= columnEnd; k++ {
		i, j = 1, len(plants)-1
		for i <= j {
			mid := (i + j) / 2
			if target < plants[mid][k] {
				j = mid - 1
			} else if target > plants[mid][k] {
				i = mid + 1
			} else {
				return true
			}
		}
	}

	return false
}

func TestFindTargetIn2DPlants(t *testing.T) {
	testCases := []struct {
		name   string
		plains [][]int
		target int
		exists bool
	}{
		{
			name: "存在",
			plains: [][]int{
				{2, 3, 6, 8},
				{4, 5, 8, 9},
				{5, 9, 10, 12},
			},
			target: 8,
			exists: true,
		},
		{
			name: "不存在",
			plains: [][]int{
				{1, 3, 5},
				{2, 5, 7},
			},
			target: 4,
			exists: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := findTargetIn2DPlants(tc.plains, tc.target)
			assert.Equal(t, tc.exists, res)
		})
	}
}
