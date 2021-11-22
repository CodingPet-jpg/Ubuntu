package main

import (
	"fmt"
	"strings"
)

// 检测字符串的同文异构,即具有相同的字符但是顺序不同
func main() {
	src := "dsa"
	des := "dsasd"
	b := judge(src, des)
	fmt.Println(b)
}

func judge(src string, des string) bool {
	if src == des {
		return false
	}
	for _, ch := range src {
		if !strings.ContainsRune(des, rune(ch)) {
			return false
		}
	}
	return true
}
