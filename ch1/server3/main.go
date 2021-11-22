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
	fmt.Fprintf(resp, "%s %s %s\n", req.Method, req.RequestURI, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(resp, "Host = %q\n", req.Host)
	fmt.Fprintf(resp, "RemoteAddr = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(resp, "Form[%q] = %q\n", k, v)
	}
	//fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
}
