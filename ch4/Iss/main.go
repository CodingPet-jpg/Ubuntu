package main

import (
	"fmt"
	"unicode"
)

func main() {
	TabTest('\n')
	s := `Hello  世界`
	bs := []byte(s)
	p := replace([]rune(s))

	fmt.Println(bs)
	fmt.Println(p)
	fmt.Printf("%s\n", string(p))
}

func replace(ru []rune) []rune {
	for i, r := range ru {
		if unicode.IsSpace(r) {
			r = rune(' ')
		}
		ru[i] = r
	}
	return ru
}

func TabTest(r rune) {
	fmt.Println(unicode.IsSpace(r))
	fmt.Println(r)
}
