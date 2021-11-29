package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var graph = make(map[string]map[string]bool)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		text := Scanner.Text()
		stringS := strings.Split(text, " ")
		from := stringS[0]
		to := stringS[1]
		addEdge(from, to)
	}
	var flag = hasEdge("from", "to")
	fmt.Println(flag)
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
