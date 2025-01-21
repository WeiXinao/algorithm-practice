package main

import "fmt"

func fileCombination(target int) [][]int {
	res := make([][]int, 0, 10)
	fileTotal := target/2 + 1
	files := make([]int, fileTotal+1)
	for i := 1; i <= fileTotal; i++ {
		files[i] = i
	}

	sum := 0
	for i, j := 1, 2; i < j && j <= fileTotal; {
		sum = (i + j) * (j - i + 1) / 2
		if sum < target {
			j++
		} else if sum > target {
			i++
		} else {
			res = append(res, files[i:j+1])
			i++
		}
	}
	return res
}

func outputMultiArr(res [][]int) {
	fmt.Print("[")
	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("[")
		for j := 0; j < len(res[i]); j++ {
			if j == 0 {
				fmt.Print(res[i][j])
			} else {
				fmt.Print(",", res[i][j])
			}
		}
		fmt.Print("]")
	}
	fmt.Print("]")
}

func main() {
	target := 0
	fmt.Print("target = ")
	fmt.Scanln(&target)
	combination := fileCombination(target)
	outputMultiArr(combination)
}
