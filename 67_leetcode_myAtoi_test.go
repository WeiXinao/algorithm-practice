package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func myAtoi(s string) int {
	res := int64(0)
  bytes := []byte(s)
	symbol := 1
	start := 0
	for _, v := range bytes {
		if v == ' ' && start == 0 {
			continue
		}	
		if v == '-' && start == 0 {
			symbol = -1	
			start = 1
			continue
		}
		if v == '+' && start == 0 {
			start = 1
			continue
		}
		if v == '0' && res == 0 {
			start = 1
			continue
		}
		if (v > '9' || v < '0') && v != ' ' && v != '+' && v != '-'  || 
			start == 1 && (v > '9' || v < '0') {  
			break
		}
	
		if res * 10 + int64(v - '0') >= - math.MinInt32 {
			res = -math.MinInt32	
			break
		} else {
			res = res * 10 + int64(v - '0')
		}
		
		if res != 0 {
			start = 1
		}
	}

	res = res * int64(symbol)
	upperLimit := math.MaxInt32
	lowerLimit := math.MinInt32
	if res > int64(upperLimit) {
		res = int64(upperLimit)
	}
	if res < int64(lowerLimit) {
		res = int64(lowerLimit)
	}
	return int(res)
}

func TestMyAtoi(t *testing.T) {
	testCases := []struct{
		name string
		s string
		expected int
	}{
		{
			name: "42",	
			s: "42",
			expected: 42,
		},
		{
			name: "-042",
			s: "-042",
			expected: -42,
		},
		{
			name: "1337c0d3",
			s: "1337c0d3",
			expected: 1337,
		},
		{
			name: "0-1",
			s: "0-1",
			expected: 0,
		},
		{
			name: "words and 987",
			s: "words and 987",
			expected: 0,
		},
		{
			name: "21474836460",
			s: "21474836460",
			expected: 2147483647,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			res := myAtoi(v.s)
			assert.Equal(t, v.expected, res)
		})	
	}
}