package main

import "fmt"

// 		  48	10
// [2, 4, 6, 8, 10 1]

func statisticalResult(arrayA []int) []int {
	if len(arrayA) <= 0 {
		return []int{}
	}
	resSlice := make([]int, len(arrayA))
	resAll, resAllWithoutZero := 1, 1
	cnt := 0
	indexOfZero := -1
	for i, val := range arrayA {
		resAll *= val
		if val == 0 {
			cnt++
			indexOfZero = i
			continue
		}
		resAllWithoutZero *= val
	}
	if cnt >= 2 {
		return resSlice
	}
	if cnt == 1 {
		resSlice[indexOfZero] = resAllWithoutZero
		return resSlice
	}
	for i := 0; i < len(arrayA); i++ {
		resSlice[i] = resAll / arrayA[i]
	}
	return resSlice
}

func main() {
	arrayA := make([]int, 0, 10)
	input := 0
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		arrayA = append(arrayA, input)
	}
	resultSlice := statisticalResult(arrayA)
	fmt.Print("[")
	for i, val := range resultSlice {
		fmt.Print(val)
		if i == len(resultSlice)-1 {
			fmt.Print("]")
		} else {
			fmt.Print(",")
		}
	}
	if len(resultSlice) <= 0 {
		fmt.Print("]")
	}
}
