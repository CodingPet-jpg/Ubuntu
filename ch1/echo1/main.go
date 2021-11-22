package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string                   // 声明两个string类型的变量，默认初始化为空字符串
	for i := 1; i < len(os.Args); i++ { // :=是短变量声明表达式
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// fmt.Println(os.Args[1:]) // TODO: 为什么打印内容带中括号
	// 因为Println格式化打印一个slice时是在打印数组,所以两边带中括号
	// 任何slice都可以以这种形式输出
}
