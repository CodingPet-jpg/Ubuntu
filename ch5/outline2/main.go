package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	fmt.Printf("Start\n")
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Printf("outline2: %v\n", err)
	}
	outline(doc, pre, post)
}

func outline(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(c, pre, post)
	}
	if post != nil {
		post(node)
	}
}

func pre(node *html.Node) {
	if node.Type == html.ElementNode {
		if node.FirstChild == nil {
			fmt.Printf("%*s<%s/>\n", depth*2, "", node.Data)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		}
		depth++
	}
}

func post(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		if node.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
		}
	}
}
