package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 123
	fmt.Println(strconv.FormatInt(int64(i), 16))
}
