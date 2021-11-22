package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "43227983471509323"
	ns := comma(s)
	fmt.Println(ns)
}

// comma接收一个string返回一个每三位插入逗号的新字符串
func comma(s string) string {
	var buf bytes.Buffer
	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		if k == 3 {
			k = 0
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		k++
	}

	var bfn bytes.Buffer
	pop(&buf, &bfn)
	return bfn.String()
}

func pop(bfo *bytes.Buffer, bfn *bytes.Buffer) {
	if bfo.Len() == 0 {
		return
	}
	b, _ := bfo.ReadByte()
	pop(bfo, bfn)
	bfn.WriteByte(b)
}
