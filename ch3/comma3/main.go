package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "-3432423531.40324"
	fmt.Println(feed(s))
}

func feed(s string) string {
	prefix := ""

	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+"){
		prefix = string(s[0])
		s = s[1:]
	}

	rs := s

	if index := strings.LastIndex(s, "."); index > 0 {
		suffix := s[index:]
		s = s[:index]
		ns := comma(s)
		rs = prefix + ns + suffix
	}
	return rs
}

func comma(s string) string {
	var spilt byte = ','
	var buf bytes.Buffer
	for i, length := 0, len(s); i < length; i++ {
		if i != 0 && (length-i)%3 == 0 {
			buf.WriteByte(spilt)
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
