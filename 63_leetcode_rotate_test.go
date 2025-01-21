package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	queue := make([]int, n)
	head, tail := 0, 0
	for i := 0; i < len(matrix)/2+len(matrix)%2; i++ {
		for j := i; j < n-i-1; j++ {
			queue[tail] = matrix[i][j]
			tail = (tail + 1) % n
		}

		for j := i; j < n-i-1; j++ {
			queue[tail] = matrix[j][n-i-1]
			tail = (tail + 1) % n
			matrix[j][n-i-1] = queue[head]
			head = (head + 1) % n
		}

		for j := n - i - 1; j >= i+1; j-- {
			queue[tail] = matrix[n-i-1][j]
			tail = (tail + 1) % n
			matrix[n-i-1][j] = queue[head]
			head = (head + 1) % n
		}

		for j := n - i - 1; j >= i+1; j-- {
			queue[tail] = matrix[j][i]
			tail = (tail + 1) % n
			matrix[j][i] = queue[head]
			head = (head + 1) % n
		}

		for j := i; j < n-i-1; j++ {
			matrix[i][j] = queue[head]
			head = (head + 1) % n
		}
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		expect [][]int
	}{
		{
			name: "三阶",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expect: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.matrix)
			assert.Equal(t, tc.expect, tc.matrix)
		})
	}
}
