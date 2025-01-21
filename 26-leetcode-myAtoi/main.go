package main

import (
	"fmt"
	"math"
)

func main() {
	str := ""
	fmt.Scanln(&str)
	res := myAtoi(str)
	fmt.Println(res)
}

func myAtoi(str string) int {
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			break
		}
		i++
	}
	str = str[i:]
	if len(str) <= 0 {
		return 0
	}
	if str[0] != '+' && str[0] != '-' && (str[0] < '0' || str[0] > '9') {
		return 0
	}
	i = len(str) - 1
	for {
		if str[i] == '+' || str[i] == '-' || str[i] >= '0' && str[i] <= '9' {
			break
		}
		i--
	}
	str = str[:i+1]

	i = 0
	for {
		if i >= len(str) {
			break
		}
		if i == 0 && str[i] != '+' && str[i] != '-' && (str[i] < '0' || str[i] > '9') {
			break
		}
		if i != 0 && (str[i] < '0' || str[i] > '9') {
			break
		}
		i++
	}
	str = str[:i]

	var res float64 = 0
	for i = len(str) - 1; i >= 0; i-- {
		if i > 0 && i < len(str) && (str[i] < '0' || str[i] > '9') {
			break
		}
		if str[i] >= '0' && str[i] <= '9' {
			res += float64(str[i]-'0') * math.Pow(float64(10), float64(len(str)-1-i))
		} else if str[i] == '+' {
			res *= 1
		} else if str[i] == '-' {
			res *= -1
		}
	}
	if res >= math.MaxInt32 {
		return math.MaxInt32
	}
	if res <= math.MinInt32 {
		return math.MinInt32
	}
	return int(res)
}
