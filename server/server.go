package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Serve(done chan bool) {
	fmt.Println("Starting server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	})
	server := &http.Server{Addr: ":3000"}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Failed to launch server with error %s", err.Error())
		}
	}()
	select {
	case <-done:
		// Shutdown gracefully, but don't wait more than 5 seconds
		ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
		if err != nil {
			// Couldn't shutdown gracefully -- forcefully cancel all listeners
			fmt.Println("Failed to gracefully shutdown", err)
			server.Close()
		}
		server.Shutdown(ctx)
	}
}
