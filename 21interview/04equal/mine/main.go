package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1, s2 := "fucc", "fucu"

	fmt.Print(equal(s1, s2))
}

func equal(s1, s2 string) bool {
	map1, map2 := make(map[string]int), make(map[string]int)

	r1, r2 := []rune(s1), []rune(s2)

	for _, v := range r1 {
		map1[string(v)]++
	}

	for _, v := range r2 {
		map2[string(v)]++
	}

	if reflect.DeepEqual(map1, map2) {
		return true
	}
	return false
}
