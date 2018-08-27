package cli

import (
	"fmt"
)

type commandInfo struct {
	usage       string
	description string
}

// Mapping of the command name to it's usage & description messages.
var commands = map[Command]commandInfo{
	Help:     commandInfo{"help", "see usage & description for all available commandInfos"},
	Man:      commandInfo{"man <COMMAND>", "see usage & description for a single commandInfo"},
	Clear:    commandInfo{"clear", "clear the terminal screen"},
	Quit:     commandInfo{"quit", "stop the server"},
	Exit:     commandInfo{"exit <WEB ADDR>", "make a HTTP GET request to an external site"},
	Announce: commandInfo{"announce <NODE ADDR>", "announce availability to peer, request peer list"},
	Relay:    commandInfo{"relay <WEB ADDR> <RELAY ADDR>", "route a request through a peer"},
}

// Handlers for various commandline functions.

func quitHandler(input []string, done chan bool) {
	fmt.Println("Stopping the server...")
	close(done)
}

func exitHandler(input []string, done chan bool) {
	if len(input) < 2 {
		fmt.Println("Not enough arguments.")
		command := commands[Exit]
		printCommand(command.usage, command.description)
		return
	}
	fmt.Println("Making HTTP GET request to ", input[1])
}

func relayHandler(input []string, done chan bool) {
	if len(input) < 3 {
		fmt.Println("Not enough arguments.")
		command := commands[Relay]
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
	commandName := commandNames[input[1]]
	command := commands[commandName]
	printCommand(command.usage, command.description)
}

func clearHandler(input []string, done chan bool) {
	fmt.Print("\033[H\033[2J")
}

// Pretty print a command name & description.
func printCommand(command string, description string) {
	fmt.Println("  \033[1m" + command + "\033[0m")
	fmt.Println("\t" + description)
}
