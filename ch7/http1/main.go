package main

import (
	"fmt"
	"net/http"
)

type database map[string]int

func (db database) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(resp, "%s:%d\n", item, price)
	}
}

func main() {
	data := database{"shoes": 10, "shirt": 5}
	http.ListenAndServe("localhost:8080", data)
}
