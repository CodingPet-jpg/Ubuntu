package main

import (
	"fmt"
)

// 模仿unix系统将路径中所有类似目录的前缀删除，同时将所有类似后缀的也删除
func main() {
	s := "/home/jokoi/GoProJ/src/jokoi.com/ch3/basename1/main.go`"
	base := basename(s)
	fmt.Println(base)
}

func basename(s string) string {
	// 将最后一个/和其之前的部分全部删除
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
  // 将最后一个.和其之后的部分全部删除
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == '.' {
			s = s[:j]
			break
		}
	}
	return s
}
