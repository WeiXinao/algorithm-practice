package main

import "fmt"

func main() {
	//a := "the sky is blue"
	//a := "  hello world!  "
	//a := "a good   example"
	a := " "
	//a = trimSpace(a)
	//fmt.Printf("%q\n", a)
	//space := splitSpace(a)
	//fmt.Printf("%#v\n", space)
	res := reverseMessage(a)
	fmt.Printf("%q\n", res)
}

func reverseMessage(message string) string {
	message = trimSpace(message)
	if len(message) == 0 {
		return ""
	}
	msgSlice := splitSpace(message)
	res := ""
	for i := len(msgSlice) - 1; i >= 0; i-- {
		res += msgSlice[i]
		if i > 0 {
			res += " "
		}
	}
	return res
}

func splitSpace(str string) []string {
	i, j := 0, 0
	strs := make([]string, 0, 10)
	for {
		for str[i] == ' ' && i < len(str) {
			i++
		}
		if j >= len(str) {
			strs = append(strs, str[i:j])
			break
		}
		if str[j] == ' ' && j > i {
			strs = append(strs, str[i:j])
			i = j
		}
		j++
	}
	return strs
}

func trimSpace(message string) string {
	msgBytes := []byte(message)
	i := 0
	for {
		if i >= len(msgBytes) || msgBytes[i] != ' ' {
			break
		}
		i++
	}
	msgBytes = msgBytes[i:]

	i = len(msgBytes) - 1
	for {
		if i <= 0 || msgBytes[i] != ' ' {
			break
		}
		i--
	}
	msgBytes = msgBytes[:i+1]

	return string(msgBytes)
}
