package main

import (
	"fmt"
)

var month []string = []string{1: "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", 12: "Dec"}

func main() {
	slice1 := month[4:9]
	slice2 := month[6:11]
	fmt.Printf("%v\t%v\n", slice1[3], slice2[1])
	fmt.Println(slice1, slice2)
	compare(slice1, slice2)
	fmt.Println("====================================")
	testlen(slice2)
	testcap(slice1)
}

func compare(s1 []string, s2 []string) {
	for i1, v1 := range s1 {
		for i2, v2 := range s2 {
			if v1 == v2 {
				fmt.Printf("index v1:%d\tv2:%d\tappears in same\t%s\n", i1, i2, v1)
			}
		}
	}
}

func testlen(s []string) {
	//fmt.Println(s[6])//可以扩容但是不能直接在当前slice访问超出自身长度的元素
	fmt.Println("未扩容前")
	fmt.Println(s)
	fmt.Println(s[:6])
	fmt.Println(len(s), cap(s))
	fmt.Println("测试切片长度越界")
	newstring := s[:7]
	fmt.Println(newstring[6])
}

func testcap(s []string) {
	fmt.Println("测试切片容量越界")
	fmt.Println(s[17])
}
