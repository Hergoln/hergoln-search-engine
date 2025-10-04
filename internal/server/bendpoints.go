package server

import (
	"io"
	"log"
	"net/http"
)

const KEY_SERVER_ADDR = "serverAddr"

func getRoot(writer http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s: Root request", ctx.Value(KEY_SERVER_ADDR))
	io.WriteString(writer, "My first http reponse in this project...")
}

func getHealthCheck(writer http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s: Healthcheck request", ctx.Value(KEY_SERVER_ADDR))
	io.WriteString(writer, "Hello, HTTP world!\n")
}

func PrepareStdMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHealthCheck)

	return mux
}
