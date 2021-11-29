package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
}
