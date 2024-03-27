package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func processFile(filePath string) (io.Reader, error){
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	stringReader := strings.NewReader(string(reader))
	return stringReader, nil
}

type Link struct {
	Href string
	Text string
}

func buildLink(node *html.Node) (Link, error){
	if node.Type != html.ElementNode && node.Data != "a" {
		return Link{}, errors.New("node is not achor tag")
	}

	var newLink Link = Link{}
	for _, attr := range(node.Attr) {
		if attr.Key == "href" {
			newLink.Href = attr.Val
		}
	}
	contentList := make([]string,0)
	contentList = extractAchorContent(node, &contentList)[1:]
	linkCotent := strings.Join(contentList, " ")
	newLink.Text = linkCotent
	return newLink, nil
}

// the first string of the returned slice will be the content of the node passsed into this func
func extractAchorContent(node *html.Node, content *[]string) ([]string){
	*content = append(*content, strings.TrimSpace(node.Data))
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractAchorContent(c, content)
	}
	return *content
}

func buildAchorList(node *html.Node, nodeList *[]*html.Node) []*html.Node {
	if node.Data == "a" {
		*nodeList = append(*nodeList, node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		buildAchorList(c, nodeList)
	}
	return *nodeList
}

func Parser(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	var anchorList []*html.Node = make([]*html.Node, 0)
	result := buildAchorList(doc, &anchorList)

	returnList := make([]Link, 0)
	for _, achor := range(result) {
		link, err := buildLink(achor)
		if err != nil {
			log.Fatal(err)
		}
		returnList = append(returnList, link)
	}
	return returnList, nil
}

func main() {
	htmlFilePath :=flag.String("file", "example2.html", "The html file path")
	flag.Parse()
	reader, err := processFile(*htmlFilePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var anchorList []*html.Node = make([]*html.Node, 0)
	result := buildAchorList(doc, &anchorList)

	for _, achor := range(result) {
		link, err := buildLink(achor)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(link)
	}

}