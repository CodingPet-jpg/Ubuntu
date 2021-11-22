package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m = make(map[string]int)
var str = "Jokoi is God"

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		text := Scanner.Text()
		strslice := strings.Split(text, " ")
		Add(strslice)
	}
	fmt.Printf("指定字串共输入%d次\n", Count(strings.Split(str, " ")))
}

// 将slice转换为字符串
func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

// 用map计数调用slice
func Add(list []string) {
	m[k(list)]++
}

// 返回slice调用map的次数
func Count(list []string) int {
	return m[k(list)]
}
