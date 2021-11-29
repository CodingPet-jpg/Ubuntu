package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // 初始化一个string:int映射的map
	// 无继承的好处就是短变量不需要写声明可以直接推测类型
	input := bufio.NewScanner(os.Stdin) // input指向输入流的buffer的头部
	for input.Scan() {
		counts[input.Text()]++ // 初始化map的key并将value加一
	}
	for line, n := range counts { // 遍历map得到两个结果分别是key和value
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) // 重复行与出现次数
		}
	}
}
