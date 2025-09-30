package main

import (
	"flag"
	"fmt"
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
	mode = flag.String("mode", "", "switches mode of operation ('server', 'client')")
	path = flag.String("path", "", "path to html file to read")

	flag.CommandLine.Usage = func() {
		fmt.Printf("This is main script of hergoln-simple-search search engine project. Right now you can either start a server mode which listens and serves on default addrs and port or run client mode which does nothing right now.\n\n")
		flag.PrintDefaults()
	}
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
		log.Printf("Setting up http 'server'...")
		RunServer()
	}

	if *mode == "client" {
		log.Printf("Setting up http 'client'...")
		RunClient()
	}
}
