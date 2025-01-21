package main

import "fmt"

func main() {
	in := 0
	putIn := make([]int, 0)
	for _, err := fmt.Scanln(&in); err == nil; _, err = fmt.Scanln(&in) {
		putIn = append(putIn, in)
	}
	fmt.Scanln(&in)
	fmt.Printf("putIn: %v\n", putIn)

	takeOut := make([]int, 0)
	for _, err := fmt.Scanln(&in); err == nil; _, err = fmt.Scanln(&in) {
		takeOut = append(takeOut, in)
	}
	fmt.Printf("takeOut: %v\n", takeOut)
	fmt.Printf("validateBookSequences(putIn, takeOut): %v\n", validateBookSequences(putIn, takeOut))
}

func validateBookSequences(putIn []int, takeOut []int) bool {
	pIn, sIn := 0, -1
	stack := make([]int, len(putIn))
	for _, v := range takeOut {
		for sIn < 0 || stack[sIn] != v && pIn < len(putIn) {
			sIn++
			stack[sIn] = putIn[pIn]
			pIn++
		}
		if stack[sIn] == v {
			sIn--
		} else {
			return false
		}
	}
	return true
}
