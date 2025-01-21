package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//func newPow(x float64, n int) float64 {
//	if n == 1 {
//		return x
//	}
//	if n == 0 {
//		return 1.0
//	}
//	if n == -1 {
//		return 1.0 / x
//	}
//	return newPow(x, n/2) * newPow(x, n-n/2)
//}
//
//func myPow(x float64, n int) float64 {
//	m := int(math.Sqrt(math.Abs(float64(n))))
//	var res float64
//	if n == 0 {
//		return 1.0
//	}
//	if n > 0 {
//		res = newPow(newPow(x, m), m) * newPow(x, n-m*m)
//	}
//	if n < 0 {
//		res = newPow(newPow(x, -m), m) * newPow(x, n+m*m)
//	}
//	return res
//}

func myPow(x float64, n int) float64 {
	var res float64
	if x == 0 {
		return 0
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	if n == 0 {
		return 1.0
	}
	m := int(math.Sqrt(math.Abs(float64(n))))
	if n > 0 {
		res = myPow(myPow(x, m), m) * myPow(x, n-m*m)
	}
	if n < 0 {
		res = myPow(myPow(x, -m), m) * myPow(x, n+m*m)
	}
	return res
}

func TestMyPow(t *testing.T) {
	testCases := []struct {
		name   string
		x      float64
		n      int
		expect string
	}{
		{
			name:   "正偶数次方",
			x:      2.00000,
			n:      10,
			expect: "1024.00000",
		},
		{
			name:   "正奇数次方",
			x:      2.10000,
			n:      3,
			expect: "9.26100",
		},
		{
			name:   "负数次方",
			x:      2.00000,
			n:      -2,
			expect: "0.25000",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := myPow(tc.x, tc.n)
			assert.Equal(t, tc.expect, fmt.Sprintf("%.5f", res))
		})
	}
}
