package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	sl = rotateRight(sl, 3)
	fmt.Println(sl)
}

func rotateLeft(sl []int, flag int) []int {
	temp := []int{}
	temp = append(temp, sl[flag:]...)
	temp = append(temp, sl[:flag]...)
	return temp
}

func rotateRight(sr []int, flag int) []int {
	temp := []int{}
	temp = append(temp, sr[len(sr)-flag:]...)
	temp = append(temp, sr[:len(sr)-flag]...)
	return temp
}
