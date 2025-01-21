package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func cuttingBamboo(bamboo_len int) int {
	max := bamboo_len - 1
	dp := make([]int, bamboo_len+1)
	for i := 1; i <= bamboo_len/2+bamboo_len%2; i++ {
		mul := cutting(i, dp) * cutting(bamboo_len-i, dp)
		if mul > max {
			max = mul
		}
	}
	if dp[bamboo_len] == 0 {
		dp[bamboo_len] = max
	}
	return max
}

func cutting(bamboo_len int, dp []int) int {
	if bamboo_len < 5 {
		return bamboo_len
	}
	if dp[bamboo_len] != 0 {
		return dp[bamboo_len]
	}
	max := bamboo_len
	for i := 1; i <= bamboo_len/2+bamboo_len%2; i++ {
		mul := cutting(i, dp) * cutting(bamboo_len-i, dp)
		if mul > max {
			max = mul
		}
	}
	if dp[bamboo_len] == 0 {
		dp[bamboo_len] = max
	}
	return max
}

func TestCuttingBamboo(t *testing.T) {
	testCases := []struct {
		name   string
		len    int
		expect int
	}{
		{
			name:   " bamboo_len = 12",
			len:    12,
			expect: 81,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := cuttingBamboo(tc.len)
			assert.Equal(t, tc.expect, res)
		})
	}
}
