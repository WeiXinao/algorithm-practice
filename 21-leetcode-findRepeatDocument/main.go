package main

import "fmt"

func main() {
	ints := inputSlice()
	outputSlice(ints)
	doc := findRepeatDocument(ints)
	fmt.Println(doc)
}

func findRepeatDocument(documents []int) int {
	docCounter := make(map[int]int)
	res := 0
	for _, doc := range documents {
		_, ok := docCounter[doc]
		if !ok {
			docCounter[doc] = 1
		} else {
			docCounter[doc]++
			res = doc
			break
		}
	}
	return res
}

func inputSlice() []int {
	input := 0
	intSlice := make([]int, 0, 10)
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		intSlice = append(intSlice, input)
	}
	return intSlice
}

func outputSlice(intSlice []int) {
	if len(intSlice) == 0 {
		fmt.Println("[]")
		return
	}
	fmt.Print("[")
	for i := 0; i < len(intSlice); i++ {
		fmt.Print(intSlice[i])
		if i == len(intSlice)-1 {
			fmt.Println("]")
		} else {
			fmt.Print(", ")
		}
	}
}
