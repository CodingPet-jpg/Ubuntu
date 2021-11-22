package main

import (
	"fmt"
)

func main() {
	unit := 4
	//不指定长度创建slicem，指定长度创建数组
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // slice的初始化方式，右边是数组
	right(i, unit)
	fmt.Println(i)
}

func left(i []int, unit int) {
	reverse(i[:unit])
	reverse(i[unit:])
	reverse(i)
}

func right(i []int, unit int) {
	reverse(i)
	reverse(i[:unit])
	reverse(i[unit:])
}

func reverse(i []int) {
	for j, k := 0, len(i)-1; j < k; j, k = j+1, k-1 {
		i[j], i[k] = i[k], i[j]
	}
}
