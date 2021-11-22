package main

// 专门处理slash的标准库是path，但是仅处理正斜杠，不分平台，更好的发挥平台是网络路径的/
// 也有专门处理文件路径的标准库path/filepath，会根据平台采取不同的切割策略
import (
	"fmt"
	"strings"
)

func main() {
	s := "/home/jokoi/GoProJ/src/jokoi.com/ch3/basename2/main.go"
	base := basename(s)
	fmt.Println(base)
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 无需做判断因为找不到slash则值取-1刚好是s[0:]即保留原完整字符串
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
