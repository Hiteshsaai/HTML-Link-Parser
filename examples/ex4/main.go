package main

import (
	"fmt"
	"strings"

	"../../link"
)

var exampleHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>`

func main() {

	////////// Parsing the HTML of the webpage
	// res, err := http.Get("https://www.golangprograms.com/golang-html-parser.html")

	// if err != nil {
	// 	panic(err)
	// }
	// // do this now so it won't be forgotten
	// defer res.Body.Close()

	// //
	// html, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// show the HTML code as a string %s
	// Converting the uint8 to string of the parsed html body
	// htmlString := string(html)
	// fmt.Println(reflect.TypeOf(string(html)))
	// fmt.Printf("%s\n", html)

	r := strings.NewReader(exampleHTML)
	// fmt.Println(r)
	links, err := link.Parser(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)

}
