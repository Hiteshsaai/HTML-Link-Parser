package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link ... This Link represents the links in the HTML tag
type Link struct {
	href string
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
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
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
