package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToUpper(str)
	letters := make(map[rune]struct{})
	for _, c := range str {
		if _, ok := letters[c]; ok {
			return false
		} else {
			letters[c] = struct{}{}
		}
	}
	return true
}

func main() {
	str := "abcd"
	fmt.Println(isUnique(str))

	str = "aA"
	fmt.Println(isUnique(str))

	str = ""
	fmt.Println(isUnique(str))

	str = "QwEtyjgdcbvN"
	fmt.Println(isUnique(str))
}
