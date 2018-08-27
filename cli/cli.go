package cli

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"sync"
)

type Command int

const (
	Help Command = iota
	Exit
	Relay
	Man
	Clear
	Quit
	Announce
)

// Commandline input causes these function handlers
// to be executed
type Handler func(argv []string, done chan bool)

var commandNames = map[string]Command{
	"help":     Help,
	"exit":     Exit,
	"relay":    Relay,
	"man":      Man,
	"clear":    Clear,
	"quit":     Quit,
	"announce": Announce,
}

// Mapping of a command name to the handler it triggers
var commandHandlers = map[Command]Handler{
	Help:  helpHandler,
	Exit:  exitHandler,
	Relay: relayHandler,
	Man:   manHandler,
	Clear: clearHandler,
	Quit:  quitHandler,
}

// Read, parse & execute command from STDIN.
// Handles safely stopping when the done channel
// closes.
func Run(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	printWelcome()
	for {
		select {
		case <-done:
			fmt.Println("Stopping CLI...")
			return
		default:
			input, err := readInput()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			handleInput(*input, done)
		}
	}
}

// Parse commandline input & route to relavant handler
func handleInput(input []string, done chan bool) {
	if len(input) == 0 {
		return
	}
	// Check if command name exists in command handler map
	commandName := commandNames[input[0]]
	if handler, ok := commandHandlers[commandName]; ok {
		// Execute handler
		handler(input, done)
	}
}

// Read a line from STDIN. Returns the input
// with leading & trailing whitespace removed.
func readInput() (*[]string, error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print(">>> ")
	input, err := buf.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "Error reading input")
	}
	inputFields := strings.Fields(input)
	return &inputFields, nil
}

// Print the initial "blurb" that appears
// when entering into the CLI.
func printWelcome() {
	fmt.Println("Onion Router 0.0.1")
	fmt.Println("[Dev Branch, Apr 1 2018]")
	fmt.Println(`Type "help" for more information, and "quit" to stop the server.`)
	fmt.Println("\033[1;31mWarning: onion router is still in development, and could be vulnerable. If you need reliable anonymity, use Tor.\033[0m")
}
