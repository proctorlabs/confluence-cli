package utility

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

//CleanHTML removes the specified information from the HTML and sets it as a string
func CleanHTML(buf []byte, bodyOnly, stripImg, cleanAdoc bool) string {
	doc, err := html.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	if bodyOnly || cleanAdoc {
		doc, err = getBody(doc)
		if err != nil {
			log.Fatal(err)
		}
	}

	if stripImg || cleanAdoc {
		doc, err = stripImgs(doc)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cleanAdoc {
		doc, err = processAdocHTML(doc)
		if err != nil {
			log.Fatal(err)
		}
	}

	result := renderNode(doc)
	return result
}

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.ToLower(n.Data) == "body" {
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
		if parent != nil && n.Type == html.ElementNode && (strings.ToLower(n.Data) == "img" || strings.ToLower(n.Data) == "script") {
			parent.RemoveChild(n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, n)
			}
		}
	}
	f(doc, nil)
	return doc, nil
}

//The purpose here is to clean and reformat an HTML file generated by ascii doctor to appear better on Confluence
//It's not perfect but it makes things better :)
func processAdocHTML(doc *html.Node) (*html.Node, error) {
	var f func(*html.Node, *html.Node)
	f = func(n, parent *html.Node) {
		if n.Type == html.ElementNode {

			//Convert <a href="#"> hash references to Confluence format nodes
			if nodeIsHashLink(n) {
				n.Data = "ac:link"
				n.Attr = []html.Attribute{html.Attribute{
					Key: "ac:anchor",
					Val: n.FirstChild.Data,
				}}
				for n.FirstChild != nil {
					n.RemoveChild(n.FirstChild)
				}
			}

			//Wrap all these elements in <strong></strong> :: h2, h3, div#toctitle
			if n.Data == "h2" || n.Data == "h3" || n.Data == "h4" || nodeHasID(n, "toctitle") {
				originalNode := n.FirstChild
				newNode := &html.Node{
					Type: html.ElementNode,
					Data: "strong",
				}
				n.InsertBefore(newNode, originalNode)
				n.RemoveChild(originalNode)
				newNode.AppendChild(originalNode)
			}

			//Insert line breaks after these elements
			if nodeHasID(n, "toc") || nodeHasClass(n, "sect1") || nodeHasClass(n, "sect2") || nodeHasClass(n, "sect3") || nodeHasClass(n, "sect4") || nodeHasClass(n, "sect5") {
				newNode := &html.Node{
					Type: html.ElementNode,
					Data: "br",
				}
				parent.InsertBefore(newNode, n)
			}

			//Scrub data that Confluence doesn't care about but can cause issues
			scrubNode(n)
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
	result := buf.String()
	return result
}

func nodeHasID(n *html.Node, id string) bool {
	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == "id" {
			return strings.ToLower(attr.Val) == strings.ToLower(id)
		}
	}
	return false
}

func nodeHasClass(n *html.Node, class string) bool {
	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == "class" {
			for _, cl := range strings.Fields(attr.Val) {
				if strings.ToLower(cl) == strings.ToLower(class) {
					return true
				}
			}
		}
	}
	return false
}

func nodeIsHashLink(n *html.Node) bool {
	if strings.ToLower(n.Data) != "a" {
		return false
	}
	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == "href" {
			return strings.HasPrefix(attr.Val, "#")
		}
	}
	return false
}

func scrubNode(n *html.Node) {
	if strings.HasPrefix(n.Data, "ac:") || n.Data == "code" {
		return
	}
	if n.Data == "div" && n.FirstChild == nil {
		n.Data = "br"
	}
	attrs := []html.Attribute{}
	for _, attr := range n.Attr {
		switch strings.ToLower(attr.Key) {
		case "id", "class", "style":
			break
		default:
			attrs = append(attrs, attr)
		}
	}
	n.Attr = attrs
}
