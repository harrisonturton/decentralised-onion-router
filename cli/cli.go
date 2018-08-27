package cli

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"sync"
)

/*
 * Functions that take commandline input, and execute
 * behaviour accordingly.
 */
type Handler func(argv []string, done chan bool)

/*
 * Mapping of command name to the handler it triggers
 */
var commandHandlers = map[string]Handler{
	"help":  helpHandler,
	"exit":  exitHandler,
	"relay": relayHandler,
	"man":   manHandler,
	"clear": clearHandler,
	"quit":  quitHandler,
}

/*
 * Handle closing goroutine when the done
 * channel closes, otherwise read, parse &
 * execute commands from STDIN.
 */
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

/*
 * Parse & execute commandline input.
 */
func handleInput(input []string, done chan bool) {
	if len(input) == 0 {
		return
	}
	// Check if command name exists in command handler map
	if handler, ok := commandHandlers[input[0]]; ok {
		// Execute handler
		handler(input, done)
	}
}

/*
 * Read a line from STDIN. Removes the newline
 * at the end. Splits the input by the spaces.
 */
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

func printWelcome() {
	fmt.Println("Onion Router 0.0.1")
	fmt.Println("[Dev Branch, Apr 1 2018]")
	fmt.Println(`Type "help" for more information, and "quit" to stop the server.`)
	fmt.Println("\033[1;31mWarning: onion router is still in development, and could be vulnerable. If you need reliable anonymity, use Tor.\033[0m")
}
