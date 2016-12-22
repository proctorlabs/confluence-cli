package utility

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func stripImgs(doc *html.Node) (*html.Node, error) {
	var f func(*html.Node, *html.Node)
	f = func(n, parent *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			parent.RemoveChild(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, n)
		}
	}
	f(doc, nil)
	return doc, nil
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

//StripHTML removes the specified information from the HTML and sets it as a string
func StripHTML(buf []byte, bodyOnly, stripImg bool) string {
	doc, err := html.Parse(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}
	if bodyOnly {
		doc, err = getBody(doc)
		if err != nil {
			panic(err)
		}
	}
	if stripImg {
		doc, err = stripImgs(doc)
		if err != nil {
			panic(err)
		}
	}
	result := renderNode(doc)
	fmt.Println(result)
	return result
}
