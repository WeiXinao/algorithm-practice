package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mechanicalAccumulator(target int) int {
	if target == 0 {
		return 0
	}
	return target + mechanicalAccumulator(target-1)
}

func TestMechanicalAccumulator(t *testing.T) {
	testCases := []struct {
		name   string
		target int
		expect int
	}{
		{
			name:   "target = 5",
			target: 5,
			expect: 15,
		},
		{
			name:   "target = 7",
			target: 7,
			expect: 28,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := mechanicalAccumulator(tc.target)
			assert.Equal(t, tc.expect, sum)
		})
	}
}
