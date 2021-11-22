package main

import "fmt"

func main() {
	is := []int{1, 2, 3, 4, 5}
	aft := remove2(is, 2)
	fmt.Println(aft)
}

func pop(i []int) int {
	temp := i[len(i)-1]
	i = i[:len(i)-1]
	return temp
}

func push(i []int, v int) []int {
	i = append(i, v)
	return i
}

func remove(i []int, index int) []int { //维持所有元素的顺序
	copy(i[index:], i[index+1:])
	return i[:len(i)-1]
}

func remove2(i []int, index int) []int {
	i[index] = i[len(i)-1]
	return i[:len(i)-1]
}
