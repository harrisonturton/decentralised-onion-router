package main

import (
	"fmt"
	"github.com/pkg/errors"
	"onion-router/server"
	"onion-router/config"
	"os"
)

const CONFIG_FILENAME = "config.yaml"

func main() {
	fmt.Println("Starting decentralised onion router...")

	fmt.Println("Reading config...")
	config, err := config.ReadConfig(CONFIG_FILENAME)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Could not read config file").Error())
		os.Exit(1)
	}
	fmt.Println("------")
	fmt.Println(config)
	fmt.Println("------")
	server.Serve(config.Server)
}
