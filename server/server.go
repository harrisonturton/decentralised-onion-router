package server

import (
	"github.com/pkg/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"onion-router/comm"
	"onion-router/exit"
)

type Config struct {
	Host string
	Port string
}

/*
 * Serve handles all the requests and routes them
 * through to HandleConnection() on a seperate]
 * goroutine.
 */
func Serve(config Config) {
	fmt.Println("Starting server...")
	fmt.Println("-------------------------------")
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleConnection)
	srv := &http.Server{
		Addr: config.Host + ":" + config.Port,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

/*
 * HandleConnection() determines whether we treat a request
 * as a relay or exit node, and then routes to HandleAsRelay()
 * or HandleAsExit() respectively.
 */
func HandleConnection(w http.ResponseWriter, req *http.Request) {
	message, err := UnmarshalRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	/* If exit node ... */
	if message.ExitAddress != nil {
		HandleAsExit(*message, w, req)
		fmt.Println("Exit node!")
		return
	}

	/* If relay node ... */
	if message.Next != nil {
		HandleAsRelay(*message, w, req)
		fmt.Println("Relay node!")
		return
	}

	/* Don't know what we're supposed to do. Return error message. */
	w.Write([]byte("Unknown message. Are you trying to send a message to a relay node, or an exit node?"))
	fmt.Println("Unknown message. Are you trying to send a message to a relay node, or an exit node?")
}

/*
 * HandleAsExit() makes a request to an external address, and passes the
 * the response back to our requestee.
 */
func HandleAsExit(message comm.Message, w http.ResponseWriter, req *http.Request) {
	exitResp, err := exit.Handle(comm.ExitMessage{
		Address: *message.ExitAddress,
		Payload: message.Payload,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte(exitResp.Address + "\n"))
	w.Write([]byte(exitResp.Payload + "\n"))
}

/*
 * HandleAsRelay() passes the request payload onto the next
 * node, and then passes the response back to our requestee.
 */
func HandleAsRelay(message comm.Message, w http.ResponseWriter, req *http.Request) {
	relayMessage := comm.RelayMessage{
		Next:    *message.Next,
		Payload: message.Payload,
	}
	w.Write([]byte(relayMessage.Next + "\n"))
}

/*
 * UnmarshalRequest() converts a raw http.Request into a nice
 * Message object.
 */
func UnmarshalRequest(req *http.Request) (*comm.Message, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read request body")
	}

	var message comm.Message
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, errors.Wrap(err, "Failed to parse request body")
	}

	return &message, nil
}

