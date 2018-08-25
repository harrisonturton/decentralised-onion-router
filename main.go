package main

import (
	"fmt"
	"partch-onion-router/cli"
	"partch-onion-router/server"
)

func main() {
	// Used to synchronse stopping
	done := make(chan bool)
	go cli.Run(done)
	go server.Serve(done)
	<-done
	fmt.Println("Goodbye!")
}
