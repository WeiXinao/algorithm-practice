package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func inventoryManagement(stock []int) int {
	length := len(stock)
	if length == 0 {
		return 0
	}
	minEle := stock[0]
	for i := 1; i < length; i++ {
		if stock[i] < minEle {
			minEle = stock[i]
		}
	}
	return minEle
}

func TestInventoryManagement(t *testing.T) {
	testCases := []struct {
		name         string
		stock        []int
		expectOutput int
	}{
		{
			name:         "[4,5,8,3,4]",
			stock:        []int{4, 5, 8, 3, 4},
			expectOutput: 3,
		},
		{
			name:         "[5,7,9,1,2]",
			stock:        []int{5, 7, 9, 1, 2},
			expectOutput: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := inventoryManagement(tc.stock)
			assert.Equal(t, tc.expectOutput, res)
		})
	}
}
