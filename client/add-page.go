package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

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

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, filepath string, bodyOnly bool) {
	page := newPage(title, spaceKey)
	response := &ConfluencePage{}
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	if bodyOnly == false {
		page.Body.Storage.Value = string(buf)
	} else {
		doc, err := html.Parse(bytes.NewReader(buf))
		if err != nil {
			panic(err)
		}
		bn, err := getBody(doc)
		if err != nil {
			panic(err)
		}
		page.Body.Storage.Value = renderNode(bn)
		fmt.Println(page.Body.Storage.Value)
	}
	c.doRequest("POST", "/rest/api/content/", page, response)
	fmt.Println("ConfluencePage Object Response", response)
}
