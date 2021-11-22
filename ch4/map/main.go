package main

import (
	"fmt"
	"sort"
)

func main() {
	Name2Age := map[string]int{"sadfasdfasf": 235234, "Bob": 12, "Jokoi": 24, "Yang": 33, "Alice": 44, "KC": 4}
	sli := Map2Slice(Name2Age)
	sort.Strings(sli)
	fmt.Println(sli)
	fmt.Println(Name2Age)
}

func Map2Slice(ages map[string]int) []string {
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	return names
}
