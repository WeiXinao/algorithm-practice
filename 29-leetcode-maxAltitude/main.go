package main

import (
	"fmt"
	"math"
)

func main() {
	input := 0
	heights := make([]int, 0)
	for _, err := fmt.Scanln(&input); err == nil; _, err = fmt.Scanln(&input) {
		heights = append(heights, input)
	}
	fmt.Scanln(&input)
	limit := 0
	fmt.Scanln(&limit)
	fmt.Printf("maxAltitude(heights, limit): %v\n", maxAltitude(heights, limit))
}

func maxAltitude(heights []int, limit int) []int {
	if len(heights) <= 0 {
		return []int{}
	}
	res := make([]int, len(heights)-limit+1)
	maxIndex := -1
	max := math.MinInt
	for i, j := 0, limit-1; j < len(heights); i, j = i+1, j+1 {
		if maxIndex < i {
			maxIndex = -1
			max = math.MinInt
			for k := i; k <= j; k++ {
				if heights[k] > max {
					max = heights[k]
					maxIndex = k
				}
			}
		} else {
			if heights[j] > max {
				max = heights[j]
				maxIndex = j
			}
		}
		res[i] = max
	}
	return res
}
