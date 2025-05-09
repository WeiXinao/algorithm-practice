package test

// import "strings"

// func isPalindrome(s string) bool {
// 	bytes := []byte(strings.ToLower(s))
// 	j := 0
// 	for i := 0; i < len(bytes); i++ {
// 		if  '0' <= bytes[i] && bytes[i] <= '9' || 
// 		'a' <= bytes[i] && bytes[i] <= 'z' ||
// 		'A' <= bytes[i] && bytes[i] <= 'Z' {
// 			bytes[j] = bytes[i] 
// 			j++
// 		}
// 	}
// 	bytes = bytes[:j]

// 	length := len(bytes)
// 	for i := 0; i < length / 2; i++ {
// 		if bytes[i] != bytes[length - 1 - i] {
// 			return false
// 		}
// 	}
// 	return true
// }