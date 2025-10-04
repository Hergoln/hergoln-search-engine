package main

import (
	"context"
	"errors"
	"hergoln-search-engine/internal/server"
	"log"
	"net"
	"net/http"
)

const (
	DEFAULT_PORT = "7"
	LOCAL_ADDR   = "127.0.0.1"
	PUBLIC_PORT  = "3333"
	ADMIN_PORT   = "4444"
)

func runPublicServer(baseCtx context.Context, cancelCtx context.CancelFunc) {
	mux := server.PrepareStdMux()
	addr := LOCAL_ADDR + ":" + PUBLIC_PORT

	publicServer := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			baseCtx = context.WithValue(baseCtx, server.KEY_SERVER_ADDR, listener.Addr().String())
			return baseCtx
		},
	}

	go func() {
		log.Printf("Starting public server at '%s' address ...\n", addr)
		err := publicServer.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed.")
			return
		} else if err != nil {
			log.Printf("Socket could not be created at '%s' host\n%s", addr, err)
			return
		}

		cancelCtx()
	}()
}

func runAdminServer(baseCtx context.Context, cancelCtx context.CancelFunc) {
	mux := server.PrepareStdMux()
	addr := LOCAL_ADDR + ":" + ADMIN_PORT

	adminServer := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			baseCtx = context.WithValue(baseCtx, server.KEY_SERVER_ADDR, listener.Addr().String())
			return baseCtx
		},
	}

	go func() {
		log.Printf("Starting admin server at '%s' address ...\n", addr)
		err := adminServer.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed.")
			return
		} else if err != nil {
			log.Printf("Socket could not be created at '%s' host\n%s", addr, err)
			return
		}

		cancelCtx()
	}()
}

func RunServer() {
	ctx, cancelCtx := context.WithCancel(context.Background()) // Background context is empty parent context

	runPublicServer(ctx, cancelCtx)
	runAdminServer(ctx, cancelCtx)

	<-ctx.Done()
}
