package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"a", "b", "c", "b", "c", "a", "b", "a"}
	map1 := count1(slice1)

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}

func count1(slice1 []string) map[string]int {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}
	return m
}
