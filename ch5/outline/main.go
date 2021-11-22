package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Printf("outline:%v:\n", err)
	}
	outline(nil, doc)
}

func outline(lists []string, node *html.Node) {
	if node.Type == html.ElementNode {
		lists = append(lists, node.Data)
		fmt.Println(lists)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(lists, c)
	}
}
