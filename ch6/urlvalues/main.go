package main

import (
	"fmt"

	"jokoi.com/ch6/net/url"
)

func main() {
	m := url.Value{"key1": {"v1", "v2"}}
	fmt.Printf("First String:%v\n", m.Get("key1"))
	m.Add("key2", "v3")
	fmt.Println(m)
	m = nil
	m.Add("key3", "v3")
	fmt.Println(m)
}
