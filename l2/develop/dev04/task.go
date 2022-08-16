package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	testArray := []string{"пятка", "пЯтак", "привет", "Тяпка", "ДОРого", "огород", "дорога", "баян", "города", "баня", "Я"}
	answer := findAnagrams(&testArray)
	fmt.Println(*answer)
}

func findAnagrams(words *[]string) *map[string]*[]string {
	result := make(map[string]*[]string)
	lettersInWords := make([]map[byte]int, 0)
	anagrams := make([][]string, 0)
	for _, v := range *words {
		v = strings.ToLower(v)
		letters := make(map[byte]int)
		for _, letter := range v {
			letters[byte(letter)]++
		}

		newAnagram := true
		for i, let := range lettersInWords {
			if reflect.DeepEqual(let, letters) {
				anagrams[i] = append(anagrams[i], v)
				newAnagram = false
				break
			}
		}
		if newAnagram {
			lettersInWords = append(lettersInWords, letters)
			anagrams = append(anagrams, []string{v})
		}
	}
	for i, v := range anagrams {
		if len(v) > 1 {
			firstWord := v[0]
			sort.Strings(anagrams[i])
			result[firstWord] = &anagrams[i]
		}
	}
	return &result
}
