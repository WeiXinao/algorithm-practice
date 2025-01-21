package test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// 1 2 3 3 3 6
// 2 1 6 6 1 2

func crackNumber(ciphertext int) int {
	if ciphertext < 0 {
		return 0
	}
	cipher := []byte(strconv.Itoa(ciphertext))
	length := len(cipher)
	dp := make([]int, length)
	dp[0] = 1
	if length >= 2 {
		n, _ := strconv.Atoi(string(cipher[:2]))
		if n <= 25 {
			dp[1] = 2
		} else {
			dp[1] = 1
		}
	}
	for i := 2; i < length; i++ {
		dp[i] = dp[i-1]
		if cipher[i-1] == '0' {
			continue
		}
		n, _ := strconv.Atoi(string(cipher[i-1 : i+1]))
		if n <= 25 {
			dp[i] += dp[i-2]
		}
	}
	return dp[length-1]
}

func TestCrackNumber(t *testing.T) {
	testCases := []struct {
		name       string
		ciphertext int
		expect     int
	}{
		{
			name:       "测试1",
			ciphertext: 216612,
			expect:     6,
		},
		{
			name:       "测试2",
			ciphertext: 506,
			expect:     1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := crackNumber(tc.ciphertext)
			assert.Equal(t, tc.expect, res)
		})
	}
}
