package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var (
	path *string
)

func parseInput() {
	path = flag.String("path", "", "path to html file to read")
	flag.Parse()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	parseInput()
	fmt.Println(*path)
	data, err := os.ReadFile(*path)
	check(err)

	reader := bytes.NewReader(data)
	hrefs, _ := gatherHRefs(reader)

	fmt.Println(hrefs)
}

func gatherHRefs(reader io.Reader) ([]string, error) {
	tokenizer := html.NewTokenizer(reader)

	var refs []string
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return refs, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			for _, attr := range token.Attr {
				if attr.Key == "href" {
					refs = append(refs, attr.Val)
				}
			}
		}
	}
}
