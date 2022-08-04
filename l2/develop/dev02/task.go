package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	newString, err := StringUnpacking("jnuin6vh")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newString)
}

func StringUnpacking(s string) (string, error) {
	result := strings.Builder{}
	var prev rune
	qty := 0
	for _, c := range s {
		intRune, err := strconv.Atoi(string(c))
		//if err == nil && prev != '\\' { //there are int
		if err == nil { //there are int
			if prev == 0 {
				return "", errors.New("no Rune before int: " + strconv.Itoa(intRune))
			}
			qty = qty*10 + intRune
		} else { //there are rune
			if qty > 0 {
				for i := 0; i <= qty; i++ {
					result.WriteRune(prev)
				}
				qty = 0
			} else if prev != 0 {
				//if prev != '\\' || (prev == '\\' && c == '\\') {
				result.WriteRune(prev)
				//}
			}
			prev = c
		}
	}
	if prev != 0 {
		for i := 0; i <= qty; i++ {
			result.WriteRune(prev)
		}
	}
	return result.String(), nil
}
