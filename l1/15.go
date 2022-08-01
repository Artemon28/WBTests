package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"unsafe"
)

var justString string

//Проблема - у нас есть длинная строка, которая занимает много памяти. Когда мы делаем срез этой строки, мы не
//копируем его, а ссылаемся на изначальную строку. Таким образом в памяти хранится вся изначальная строка
// и сборщик мусора не освобождает, даже когда мы ей, не пользуемся, например если мы вернём эту строку из функции.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&v)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data)
}

func correctSomeFunc1() string {
	v := createHugeString(1 << 10)
	justString := v[:100]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&v)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data)
	copyString := make([]byte, len(justString))
	copy(copyString, justString)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&copyString)).Data)
	return string(copyString)
}

//ещё есть проблема со срезом по количеству байт, но наш string может быть массивом rune, тогда мы обрежем руну.
//Поэтому как вариант решения не копировать строку напрямую, а сделать массив рун,
//который уже будет ссылаться на другую область памяти

func correctSomeFunc2() string {
	v := createHugeString(1 << 10)
	justString := []rune(v)[:100]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&v)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&justString)).Data)
	return string(justString)
}

func createHugeString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main() {
	correctSomeFunc1()
	fmt.Println("")
	correctSomeFunc2()
}
