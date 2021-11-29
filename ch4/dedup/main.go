package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	deMap := make(map[string]bool)
	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		text := Scanner.Text()
		if !deMap[text] {
			deMap[text] = true
			fmt.Println(text)
		}
	}
	if err := Scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}
