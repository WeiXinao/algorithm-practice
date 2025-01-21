package main

import "fmt"

func trainingPlan(actions []int) []int {
	for i, j := 0, 0; j < len(actions); j++ {
		if actions[j]%2 != 0 {
			if actions[i] != actions[j] {
				actions[i], actions[j] = actions[j], actions[i]
			}
			i++
		}
	}
	return actions
}

func inputArr() []int {
	arr := make([]int, 0, 10)
	in := 0
	for {
		_, err := fmt.Scanln(&in)
		if err != nil {
			break
		}
		arr = append(arr, in)
	}
	return arr
}

func outputArr(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	arr := inputArr()
	trainingPlan(arr)
	outputArr(arr)
}
