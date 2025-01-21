package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWardrobeFinishing(t *testing.T) {
	testCases := []struct {
		name   string
		m      int
		n      int
		cnt    int
		expect int
	}{
		{
			name:   "示例1",
			m:      4,
			n:      7,
			cnt:    5,
			expect: 18,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := wardrobeFinishing(tc.m, tc.n, tc.cnt)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func wardrobeFinishing(m int, n int, cnt int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	move(grid, 0, 0, cnt)

	steps := 0
	for i := range grid {
		for j := range grid[i] {
			steps += grid[i][j]
		}
	}
	return steps
}

func move(grid [][]int, i, j, cnt int) {
	if digit(i) > cnt || digit(j) > cnt {
		return
	}
	if digit(i)+digit(j) > cnt {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	if j+1 < len(grid[i]) {
		move(grid, i, j+1, cnt)
	}
	if i+1 < len(grid) {
		move(grid, i+1, j, cnt)
	}
}

func digit(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
