package main

import "fmt"

func pathEncryption(path string) string {
	pathBytes := []byte(path)
	for i := 0; i < len(pathBytes); i++ {
		if pathBytes[i] == '.' {
			pathBytes[i] = ' '
		}
	}
	return string(pathBytes)
}

func main() {
	str := ""
	fmt.Scanln(&str)
	res := pathEncryption(str)
	fmt.Println(res)
}
