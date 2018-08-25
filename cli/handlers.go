package cli

import (
	"fmt"
)

type command struct {
	usage       string
	description string
}

var commands = map[string]command{
	"help":     command{"help", "see usage & description for all available commands"},
	"man":      command{"man <COMMAND>", "see usage & description for a single command"},
	"clear":    command{"clear", "clear the terminal screen"},
	"quit":     command{"quit", "stop the server"},
	"exit":     command{"exit <WEB ADDR>", "make a HTTP GET request to an external site"},
	"announce": command{"announce <NODE ADDR>", "announce availability to peer, request peer list"},
	"relay":    command{"relay <WEB ADDR> <RELAY ADDR>", "route a request through a peer"},
}

func quitHandler(input []string, done chan bool) {
	fmt.Println("Stopping the server...")
	close(done)
}

func exitHandler(input []string, done chan bool) {
	if len(input) < 2 {
		fmt.Println("Not enough arguments.")
		command := commands["exit"]
		printCommand(command.usage, command.description)
		return
	}
	fmt.Println("Making HTTP GET request to ", input[1])
}

func relayHandler(input []string, done chan bool) {
	if len(input) < 3 {
		fmt.Println("Not enough arguments.")
		command := commands["relay"]
		printCommand(command.usage, command.description)
		return
	}
	fmt.Printf("Relaying request to %s through node %s\n", input[1], input[2])
}

func helpHandler(input []string, done chan bool) {
	for _, command := range commands {
		printCommand(command.usage, command.description)
	}
}

func manHandler(input []string, done chan bool) {
	if len(input) != 2 {
		fmt.Println("What manual page do you want?")
		return
	}
	command := commands[input[1]]
	printCommand(command.usage, command.description)
}

func clearHandler(input []string, done chan bool) {
	fmt.Print("\033[H\033[2J")
}

/*
 * Pretty print a command name & description.
 */
func printCommand(command string, description string) {
	fmt.Println("  \033[1m" + command + "\033[0m")
	fmt.Println("\t" + description)
}
