package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/count" {
		fmt.Fprintf(resp, "count:%d", count)
		return
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
}
