package main

import (
	"onion-router/server"
	"fmt"
)

func main() {
	fmt.Println("Starting decentralised onion router...");
	server.Serve()
}
