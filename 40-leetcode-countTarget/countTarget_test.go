package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countTarget(scores []int, target int) int {
	if len(scores) == 0 {
		return 0
	}

	cnt := 0
	for _, score := range scores {
		if score == target {
			cnt++
		}
	}

	return cnt
}

func TestCountTarget(t *testing.T) {
	testCases := []struct {
		name   string
		scores []int
		target int
		output int
	}{
		{
			name:   "存在目标成绩",
			scores: []int{2, 2, 3, 4, 4, 4, 5, 6, 6, 8},
			target: 4,
			output: 3,
		},
		{
			name:   "不存在目标成绩",
			scores: []int{1, 2, 3, 5, 7, 9},
			target: 6,
			output: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := countTarget(tc.scores, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
