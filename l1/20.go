package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	reverseString := strings.Fields(str)
	for i, j := 0, len(reverseString)-1; i < j; i, j = i+1, j-1 {
		reverseString[i], reverseString[j] = reverseString[j], reverseString[i]
	}
	return strings.Join(reverseString, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
