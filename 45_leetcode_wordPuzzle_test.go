package test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWordPuzzle(t *testing.T) {
	testCases := []struct {
		name   string
		grid   [][]byte
		target string
		expect bool
	}{
		{
			name: "从第一个格子开始，存在",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target: "ABCCED",
			expect: true,
		},
		{
			name: "以中间格子开始，存在",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target: "SEE",
			expect: true,
		},
		{
			name: "不存在",
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			target: "ABCB",
			expect: false,
		},
		{
			name: "迷宫只有一个元素",
			grid: [][]byte{
				{'a'},
			},
			target: "a",
			expect: true,
		},
		{
			name: "全部走一遍",
			grid: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			target: "cdba",
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := wordPuzzle(tc.grid, tc.target)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func wordPuzzle(grid [][]byte, target string) bool {
	for i := range grid {
		for j := range grid[i] {
			if walkPuzzle(grid, "", &target, i, j) {
				return true
			}
		}
	}
	return false
}

func walkPuzzle(grid [][]byte, str string, target *string, i, j int) bool {
	if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 {
		return false
	}
	if grid[i][j] == 0 {
		return false
	}
	newGrid := make([][]byte, len(grid))
	for k := range newGrid {
		newGrid[k] = make([]byte, len(grid[k]))
		copy(newGrid[k], grid[k])
	}

	newGrid[i][j] = 0
	str += string(grid[i][j])
	if !strings.HasPrefix(*target, str) {
		return false
	}
	if str == *target {
		return true
	}
	return walkPuzzle(newGrid, str, target, i+1, j) || walkPuzzle(newGrid, str, target, i, j-1) ||
		walkPuzzle(newGrid, str, target, i-1, j) || walkPuzzle(newGrid, str, target, i, j+1)
}
