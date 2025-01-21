package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func takeAttendance(records []int) int {
	for i := range records {
		if records[i] != i {
			return i
		}
	}
	return len(records)
}

func TestTakeAttendance(t *testing.T) {
	testCases := []struct {
		name      string
		records   []int
		expectOut int
	}{
		{
			name:      "偶数缺席",
			records:   []int{0, 1, 2, 3, 5},
			expectOut: 4,
		},
		{
			name:      "奇数缺席",
			records:   []int{0, 1, 2, 3, 4, 5, 6, 8},
			expectOut: 7,
		},
		{
			name:      "最后一个缺席",
			records:   []int{0},
			expectOut: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attendance := takeAttendance(tc.records)
			assert.Equal(t, tc.expectOut, attendance)
		})
	}
}
