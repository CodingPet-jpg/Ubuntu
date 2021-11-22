package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	intArrays := []int{5, 4, 1, 7, 6}
	sorted := Sort(intArrays)
	fmt.Printf("before sort:%v\n", intArrays)
	fmt.Printf("after sort:%v\n", sorted)
}

// 传入的值调用add函数挂在树上
// 再通过appendValues函数将排序后的值传回values
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// 将树上的结果按排序插入slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 将value挂在树上
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
