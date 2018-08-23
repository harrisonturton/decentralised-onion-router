package main

import (
	"fmt"
	"onion-router/server"
)

func main() {
	fmt.Println("Starting decentralised onion router...")
	server.Serve()
}
