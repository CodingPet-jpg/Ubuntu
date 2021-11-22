package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 从命令行读取文件名参数,打开文件获取文件所有内容
	// 以换行符为切割点放入map,再讲数量大于1的重复行打印
	filelist := os.Args[1:]
	counts := make(map[string]int)
	for _, filename := range filelist {
		// ReadFile返回的是slice
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}
		// 将整块slice从换行符开始分割成字符slice的silce
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for data, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, data)
		}
	}
}
