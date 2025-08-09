package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	hrefs := gatherHRefs(data)

	fmt.Println(hrefs)
}

func gatherHRefs(data []byte) []string {
	datas := string(data)

	splits := strings.Split(datas, "href=")
	splits = splits[1:]
	for i, e := range splits {
		e = e[1:]
		e = e[:strings.Index(e, "\"")]
		splits[i] = e
	}

	return splits
}
