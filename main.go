package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	mode  *string
	path  *string
	modes = []string{"single_file", "server"}
)

func parseInput() {
	mode = flag.String("mode", "", "mode of operation")
	path = flag.String("path", "", "path to html file to read")
	flag.Parse()
}

func main() {
	parseInput()

	if !slices.Contains(modes, *mode) {
		log.Printf("Mode '%s' is not in allowed modes (%v)\n", *mode, modes)
	}

	if *mode == "single_file" {
		log.Printf("Starting 'single_file' mode, reading hrefs from file")
		runSingleFileScan()
	}

	if *mode == "server" {
		log.Printf("Starting 'server' mode.")
	}
}

func runSingleFileScan() {
	fmt.Println(*path)
	data, err := os.ReadFile(*path)
	check(err)

	reader := bytes.NewReader(data)
	hrefs, _ := gatherHRefs(reader)

	fmt.Println(hrefs)
}
