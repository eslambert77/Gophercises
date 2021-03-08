package htmlwayfind

import (
	"bytes"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

// ReadHTML parses from an HTML file and returns the root node of the HTML Tree.
func ReadHTML(filename string) *html.Node {
	htm, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(htm)

	docTree, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	return docTree
}
