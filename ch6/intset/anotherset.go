package intset

import "fmt"

func P() {
	v := Intset{words: []uint64{1, 2, 3}}
	fmt.Println(v.words)
}
