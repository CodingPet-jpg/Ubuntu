package main

import (
	"fmt"
)

//消除重复字符
func main() {
	str := "baaac"
	st := delete([]byte(str))
	fmt.Println(string(st))
}

func delete(str []byte) []byte {
	var vec byte = 0
	i := 0
	for index, s := range str {
		fmt.Printf("第%d次\n", index)
		if s != vec {
			vec = s
			str[i] = s
			i++
		}
	}
	return str[:i]
}
