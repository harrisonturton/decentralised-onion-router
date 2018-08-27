package main

import (
	"fmt"
	"onion-router/cli"
	"onion-router/server"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Used to synchronse stopping
	done := make(chan bool)

	wg.Add(2)
	go cli.Run(done, &wg)
	go server.Run(done, &wg)
	wg.Wait()

	fmt.Println("Goodbye!")
}
