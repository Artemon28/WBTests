package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//мой ручной вариант, здесь я не знаю как итерируется range по строке
func reverse(str string) string {
	var b strings.Builder
	runeString := make([]rune, utf8.RuneCountInString(str))
	for i, w, j := 0, 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		runeString[j] = runeValue
		j++
		w = width
	}
	for i := len(runeString) - 1; i >= 0; i-- {
		b.WriteRune(runeString[i])
	}
	return b.String()
}

//Вариант с выделением места под всю строку и последующее заполнение с конца, таким образом мы переставим байты
//в правильном порядке. Интересный факт: range итерируется по рунам, поэтому это работает.
func reverse2(str string) string {
	length := len(str) - 1
	reverseString := make([]rune, length)
	for _, c := range str {
		length--
		reverseString[length] = c
	}
	return string(reverseString)
}

//также можно конвертировать строку в массив рун и он правильно сложит руны. Самый быстрый алгоритм получается
func reverse3(str string) string {
	reverseString := []rune(str)
	for i, j := 0, len(reverseString)-1; i < j; i, j = i+1, j-1 {
		reverseString[i], reverseString[j] = reverseString[j], reverseString[i]
	}
	return string(reverseString)
}

func main() {
	str := "狐 jumped over the lazy 犬"
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
	fmt.Println(reverse3(str))
}
