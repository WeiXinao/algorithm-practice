package main

import "fmt"

// s3cur1tyC0d3
// uc3sr1tyC0d3
// uc3s3d0cyt1r
// r1tyc0d3s3cu

func main() {
	password := ""
	fmt.Scanln(&password)
	target := 0
	fmt.Scanln(&target)
	res := dynamicPassword(password, target)
	fmt.Println(res)
}

func dynamicPassword(password string, target int) string {
	res := rangeReverse(password, 0, target-1)
	res = rangeReverse(res, target, len(res)-1)
	res = rangeReverse(res, 0, len(res)-1)
	return res
}

func rangeReverse(str string, lo, hi int) string {
	bytes := []byte(str)
	for i := 0; i < (hi-lo+1)/2; i++ {
		bytes[lo+i], bytes[hi-i] = bytes[hi-i], bytes[lo+i]
	}
	return string(bytes)
}
