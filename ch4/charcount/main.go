package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var countMap = make(map[rune]int)
var utflen = [utf8.UTFMax + 1]int{}
var invaild = 0
var letter = 0
var digit = 0

func main() {
	CountFromStdin()
	if invaild > 0 {
		fmt.Fprintf(os.Stderr, "invaild:%d", invaild)
	}
	if digit > 0 {
		fmt.Printf("degit:%d\n", digit)
	}
	if letter > 0 {
		fmt.Printf("letter:%d\n", letter)
	}
	for index, count := range utflen {
		if index > 0 {
			fmt.Printf("The rune use %d bytes's count is %d\n", index, count)
		}
	}
	for k, v := range countMap {
		fmt.Printf("Rune:\t%4q\thas count\t %4d\n", k, v)
	}
}

func CountFromStdin() {
	Reader := bufio.NewReader(os.Stdin)
	for {
		ru, nbytes, err := Reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "err:%v\n", err)
			os.Exit(1)
		}
		if ru == unicode.ReplacementChar && nbytes == 1 {
			invaild++
			continue
		}
		if unicode.IsLetter(ru) {
			letter++
		}
		if unicode.IsDigit(ru) {
			digit++
		}
		utflen[nbytes]++
		countMap[ru]++
	}
}
