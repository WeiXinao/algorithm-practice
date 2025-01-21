package main

import (
	"fmt"
)

/*
1 2 3
8 9 4
7 6 5

1  2  3  4
12 13 14 5
11 16 15 6
10 9  8  7

0 7 0
0 9 0
0 6 0

2 3
*/

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	arrlen := len(array) * len(array[0])
	ints := make([]int, arrlen)
	cnt := 0
	for i := 0; i < len(array)/2; i++ {
		for j := i; j < len(array[0])-1-i; j++ {
			if cnt < arrlen {
				ints[cnt] = array[i][j]
				cnt++
			}
		}
		for j := i; j < len(array)-1-i; j++ {
			if cnt < arrlen {
				ints[cnt] = array[j][len(array[0])-1-i]
				cnt++
			}
		}
		for j := i; j < len(array[0])-1-i; j++ {
			if cnt < arrlen {
				ints[cnt] = array[len(array)-1-i][len(array[0])-1-j]
				cnt++
			}
		}
		for j := i; j < len(array)-1-i; j++ {
			if cnt < arrlen {
				ints[cnt] = array[len(array)-1-j][i]
				cnt++
			}
		}
	}
	if len(array)%2 != 0 {
		for j := len(array) / 2; j < len(array[0])-len(array)/2; j++ {
			if cnt < arrlen {
				ints[cnt] = array[len(array)/2][j]
				cnt++
			}
		}
	}
	return ints
}

func main() {
	//array1 := [][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}
	//array2 := [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}
	//array3 := [][]int{[]int{2, 3}}
	array4 := [][]int{[]int{7}, []int{9}, []int{6}}
	ints := spiralArray(array4)
	fmt.Print("[")
	for i, val := range ints {
		fmt.Print(val)
		if i == len(ints)-1 {
			fmt.Print("]")
		} else {
			fmt.Print(", ")
		}
	}
}
