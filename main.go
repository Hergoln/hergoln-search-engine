package main

import (
	"flag"
	"hergoln-search-engine/internal/processing"
	"log"
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
		processing.RunSingleFileScan(path)
	}

	if *mode == "server" {
		log.Printf("Starting http 'server' mode...")
		RunServer()
	}

	if *mode == "client" {
		log.Printf("Starting http 'client' mode...")
		RunClient()
	}
}
