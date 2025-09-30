package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = "7"
	DEFAULT_ADDR = "127.0.0.1"
	DEFAULT_HOST = "127.0.0.1:7"
)

func getRoot(writer http.ResponseWriter, r *http.Request) {
	log.Println("Root request")
	io.WriteString(writer, "My first http reponse in this project...")
}

func getHealthCheck(writer http.ResponseWriter, r *http.Request) {
	log.Println("Healthcheck request")
	io.WriteString(writer, "Hello, HTTP world!\n")
}

func prepareMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHealthCheck)

	return mux
}

func RunServer() {
	mux := prepareMux()
	host := DEFAULT_ADDR + ":3333"
	log.Printf("Server started at '%s' address...\n", host)
	err := http.ListenAndServe(host, mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Printf("Socket could not be created in local host at '%s'\n", DEFAULT_HOST)
		return
	}
}
