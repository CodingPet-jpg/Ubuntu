package main

import (
	"fmt"
	"strings"
)

func main() {
	s := commaForFloat("1244732559.834")
	fmt.Println(s)
	slice := string2slice()
	ss := string(slice)
	fmt.Printf("%s\n", ss)
}

// 包装器模式
func commaForFloat(s string) string {
	i := strings.LastIndex(s, ".")
	afterPoint := s[i:]
	fs := comma(s[:i])
	return fs + afterPoint
}

// TODO: 学完debug后观测添加逗号时每个帧的局部变量s会不会被修改,即这个s是引用类型还是值类型
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 字符串和字节数组可以随意转换
func string2slice() []byte {
	s := "Hello World"
	b := []byte(s)
	return b
}
