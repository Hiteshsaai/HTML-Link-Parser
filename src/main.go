package main

import (
	"fmt"
	"link"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`

func main() {

	r := strings.NewReader(exampleHtml)

	links, err := link.Parser(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v+", links)

}