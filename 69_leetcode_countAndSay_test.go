package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	builder := strings.Builder{}
	s := countAndSay(n-1)
	b := []byte(s)
	cnt := 1
	char := b[0]
	
	for i := 0; i < len(b); i++ {
		if i == 0 {
			continue
		}
		if b[i] == b[i - 1] {
			cnt++
		} else {
			builder.WriteString(fmt.Sprintf("%d%c", cnt,char))
			cnt = 1
			char = b[i]
		}
	}
	builder.WriteString(fmt.Sprintf("%d%c", cnt,char))
	return builder.String()
}

func TestCountAndSay(t *testing.T) {
	testCases := []struct{
		name string
		n int
		expect string
	} {
		{
			name:   "1",
			n:      1,
			expect: "1",
		},
		{
			name:   "2",
			n:      2,
			expect: "11",
		},
		{
			name:   "3",
			n:      3,
			expect: "21",
		},
		{
			name:   "4",
			n:      4,
			expect: "1211",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countAndSay(tc.n)
			assert.Equal(t, tc.expect, result)
		})	
	}
}