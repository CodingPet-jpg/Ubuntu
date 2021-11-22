package main

import (
	"bytes"
	"fmt"
)

//穿入{1，2，3}的数组输出字符串[1,2,3]
func main() {
	ia := []int{1, 2, 3}
	s := intsToString(ia)
	fmt.Println(s)
}

func intsToString(values []int) string {
	var buf bytes.Buffer // Buffer变量为零值无须初始化
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v) //Buffer看做文件作为接收器
	}
	buf.WriteByte(']')
	return buf.String()
}
