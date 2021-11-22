package main

import (
	"fmt"
)

func main() {
	s := []string{"Hello World", "", "", "", "GOOD JOB"}
	fmt.Printf("处理前:%q\n", s)
	newS := noneempty2(s)
	fmt.Printf("处理后:%q\n", newS)
}

func noneempty(str []string) []string {
	i := 0
	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

func noneempty2(str []string) []string {
	out := str[:0]
	for _, v := range str {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
