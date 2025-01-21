package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}

	return dp[n-1]
}

func TestNthUglyNumber(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		expect int
	}{
		{
			name:   "10",
			n:      10,
			expect: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := nthUglyNumber(tc.n)
			assert.Equal(t, tc.expect, res)
		})
	}
}
