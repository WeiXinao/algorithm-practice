package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyTreeOrder(t *testing.T) {
	testCases := []struct {
		name      string
		postorder []int
		expect    bool
	}{
		{
			name:      "不是二叉搜索树",
			postorder: []int{4, 9, 6, 5, 8},
			expect:    false,
		},
		{
			name:      "是二叉搜索树",
			postorder: []int{4, 6, 5, 9, 8},
			expect:    true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isSearchTree := verifyTreeOrder(tc.postorder)
			assert.Equal(t, tc.expect, isSearchTree)
		})
	}
}

func verifyTreeOrder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	var (
		rootIndex = len(postorder) - 1
		root      = postorder[rootIndex]
	)
	i := 0
	for postorder[i] < root {
		i++
	}
	j := i
	for postorder[j] > root {
		j++
	}
	return j == len(postorder)-1 && verifyTreeOrder(postorder[0:i]) && verifyTreeOrder(postorder[i:rootIndex])
}
