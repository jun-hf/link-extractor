package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func processDoc(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := node.FirstChild; c!= nil; c = c.NextSibling {
		processDoc(c)
	}
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	processDoc(doc)
}