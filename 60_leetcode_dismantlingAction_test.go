package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	dp[i] = dp[i - 1]
	dp[i] = dp[i - 1] + 1

	dp[0] = 1
    1 2
	d b a s c D d a d
    k = 2
	1 2 2 2 3
	p w w k e w
*/

func dismantlingAction(arr string) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}
	arrBytes := []byte(arr)
	length := len(arrBytes)
	dp := make([]int, length)
	dp[0] = 1
	k := 0
label:
	for i := 1; i < length; i++ {
		for j := 1; j <= min(dp[i-1], i-k); j++ {
			if arrBytes[i-j] == arrBytes[i] {
				k = i - j + 1
				dp[i] = dp[i-1]
				continue label
			}
		}
		if i-k+1 > dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[length-1]
}

func TestDismantlingAction(t *testing.T) {
	testCases := []struct {
		name   string
		arr    string
		expect int
	}{
		{
			name:   "pwwkew",
			arr:    "pwwkew",
			expect: 3,
		},
		{
			name:   "eeydgwdykpv",
			arr:    "eeydgwdykpv",
			expect: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := dismantlingAction(tc.arr)
			assert.Equal(t, tc.expect, res)
		})
	}
}
