package main

import (
	"fmt"
	"strings"

	"../../link"
)

var exampleHTML = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
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
