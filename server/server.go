package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"sync"
)

func Serve(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	server := MakeServer("3000")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				fmt.Printf("Failed to launch server with error %s\n", err.Error())
				log.Printf("Failed to launch server with error %s", err.Error())
			}
		}
	}()
	fmt.Println("Blocking until server quit...")
	<-done
	fmt.Println("Inside server shutdown...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Couldn't gracefully shutdown")
		log.Fatal(err)
		//server.Close()
		return
	}
	
	fmt.Println("Server shutdown gracefully.")
}

func MakeServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, "Hello World")
	})
	return &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}
}
