package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const port = 3000
const shutdownTimeout = 5 * time.Second

// Run starts the server and attempts to gracefully
// shutdown when the done channel is closed.
func Run(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	server := makeServer(port)
	go serve(server)
	<-done // Wait for done signal
	if err := attemptShutdown(server); err != nil {
		fmt.Println("Failed to gracefully shutdown.")
		return
	}
	fmt.Println("Server shutdown gracefully.")
}

// Launch the server & listen for requests on
// the specified port.
func serve(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("Failed to launch server with error %s\n", err.Error())
		}
	}
}

// Attempts to gracefully shutdown the server without
// interrupting any active connections. Waits for
// shutdownTimeout before rudely cancelling any active
// listeners & forcing the shutdown.
// See http.Shutdown & http.Cancel for more details.
func attemptShutdown(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}

// Creates a http.Server with the proper handlers
func makeServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello World")
	})
	fmt.Println(strconv.Itoa(port))
	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}
}
