package main

import (
	"fmt"
	"sort"
)

func main() {
	var set1 = make(map[string]struct{})
	set1["key1"] = struct{}{}
	set1["key2"] = struct{}{}
	set1["key3"] = struct{}{}
	var set2 = make(map[string]struct{})
	set2["key2"] = struct{}{}
	set2["key5"] = struct{}{}
	set2["key3"] = struct{}{}

	var resultSet = make(map[string]struct{})

	//решение за O(n^2) и O(n) по памяти на новое множество
	for k1, _ := range set1 {
		for k2, _ := range set2 {
			if k1 == k2 {
				resultSet[k1] = struct{}{}
				break
			}
		}
	}
	fmt.Println(resultSet)

	//решение за O(nLogn) и лишнии O(n) на массивы
	var resultSet2 = make(map[string]struct{})
	arr1 := []string{}
	for k, _ := range set1 {
		arr1 = append(arr1, k)
	}
	arr2 := []string{}
	for k, _ := range set2 {
		arr2 = append(arr2, k)
	}
	sort.Strings(arr1)
	sort.Strings(arr2)
	i1 := 0
	i2 := 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] == arr2[i2] {
			resultSet2[arr1[i1]] = struct{}{}
			i1++
			i2++
		} else if arr1[i1] < arr2[i2] {
			i1++
		} else {
			i2++
		}
	}

	fmt.Println(resultSet2)
}
