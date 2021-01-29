package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link ... This Link represents the links in the HTML tag
type Link struct {
	Href string
	Text string
}

// Parser ... Parse will take an HTML document and it will
// return slices
func Parser(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var ret []Link
	for _, node := range nodes {
		ret = append(ret, buildLink(node))
	}
	return ret, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Key
			break
		}
	}
	ret.Text = getText(n)
	return ret
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func dfs(n *html.Node, padding string) {
	if n.Type == html.ElementNode {
		fmt.Println("<" + n.Data + ">")
	} else {
		fmt.Println(padding, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, " ")
	}

}
