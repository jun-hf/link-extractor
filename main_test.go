package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestProcessFile(t *testing.T) {
    // Create a temporary test HTML file
    htmlContent := `
        <html>
            <body>
                <a href="https://www.example.com">Example</a>
                <a href="https://www.example2.com">Example 2</a>
            </body>
        </html>`
    tempFile := createTempFile(t, htmlContent)
    defer tempFile.Close()

    // Test processFile function
    reader, err := processFile(tempFile.Name())
    if err != nil {
        t.Errorf("processFile returned error: %v", err)
    }
    content := readAllString(t, reader)
    expectedContent := `<html><body><a href="https://www.example.com">Example</a><a href="https://www.example2.com">Example 2</a></body></html>`
    if content != expectedContent {
        t.Errorf("processFile content = %s; want %s", content, expectedContent)
    }
}

func TestBuildLink(t *testing.T) {
    // Create a sample <a> HTML node
    anchorHTML := `<a href="https://www.example.com">Example</a>`
    node, err := parseHTML(anchorHTML)
    if err != nil {
        t.Fatal(err)
    }

    // Test buildLink function
    link, err := buildLink(node)
    if err != nil {
        t.Errorf("buildLink returned error: %v", err)
    }
    expectedLink := Link{Href: "https://www.example.com", Text: "Example"}
    if link != expectedLink {
        t.Errorf("buildLink result = %+v; want %+v", link, expectedLink)
    }
}

// Helper functions

func createTempFile(t *testing.T, content string) *os.File {
    tempFile, err := os.CreateTemp("", "test.html")
    if err != nil {
        t.Fatal(err)
    }
    defer tempFile.Close()
    _, err = tempFile.WriteString(content)
    if err != nil {
        t.Fatal(err)
    }
    return tempFile
}

func readAllString(t *testing.T, reader io.Reader) string {
    content, err := io.ReadAll(reader)
    if err != nil {
        t.Fatal(err)
    }
    return string(content)
}

func parseHTML(htmlContent string) (*html.Node, error) {
    reader := strings.NewReader(htmlContent)
    return html.Parse(reader)
}
