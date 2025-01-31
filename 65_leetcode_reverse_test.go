package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(x int) int {
	symbol := 1
	if x < 0 {
		x *= -1
		symbol = -1
	}
	res := 0
	for x > 0 {
		res = res * 10 + (x % 10)
		x /= 10
	}
	res = res * symbol
	if res > int(math.Pow(2, 31) - 1) || res < int(-math.Pow(2, 31)) {
		return 0
	}
	return res
}

func TestReserse(t *testing.T) {
	testCases := []struct{
		name string
		x int
		expect int
	}{
		{
			name: "正数",
			x: 123,
			expect: 321,	
		},
		{
			name: "负数",
			x: -123,
			expect: -321,
		},
		{
			name: "整数，结尾为0",
			x: 120,
			expect: 21,
		},
		{
			name: "零",
			x: 0,
			expect: 0,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			res := reverse(v.x)
			assert.Equal(t, v.expect, res)
		})
	}
}