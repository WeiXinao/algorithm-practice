package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = ""
	fmt.Scanln(&s)
	fmt.Printf("validNumber(s): %v\n", validNumber(s))
}

func validNumber(s string) bool {
	s = strings.TrimSpace(s)
	i := 0
	sep := ""
	for i = 0; i < len(s); i++ {
		if s[i] == 'e' {
			sep = "e"
			break
		}
		if s[i] == 'E' {
			sep = "E"
			break
		}
	}
	if i == 0 || i == len(s)-1 {
		return false
	}
	if i == len(s) {
		return isInteger(s) || isFloat(s)
	}
	strs := strings.Split(s, sep)
	if len(strs) >= 3 {
		return false
	}
	return (isInteger(strs[0]) || isFloat(strs[0])) && isInteger(strs[1])
}

func isFloat(s string) bool {
	if s == "+" || s == "-" || s == "." {
		return false
	}
	if s == "+." || s == "-." {
		return false
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			cnt++
		}
		if i == 0 && (s[i] > '9' || s[i] < '0') && s[i] != '+' && s[i] != '-' && s[i] != '.' {
			return false
		}
		if i > 0 && (s[i] > '9' || s[i] < '0') && s[i] != '.' {
			return false
		}
	}
	if cnt != 1 {
		return false
	}
	return true
}

func isInteger(s string) bool {
	if s == "+" || s == "-" {
		return false
	}
	if (s[0] > '9' || s[0] < '0') && s[0] != '+' && s[0] != '-' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}
	}
	return true
}
