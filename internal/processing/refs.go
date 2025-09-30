package processing

import (
	"io"

	"golang.org/x/net/html"
)

func GatherHRefs(reader io.Reader) ([]string, error) {
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
