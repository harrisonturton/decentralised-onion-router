package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	/* Says if we're to operate as an exit node */
	ExitAddress *string `json:"exit_address"`

	/* If not exit node, then relay. Address of node to forward to */
	Next        *string `json:"next"`

	/* Payload to transfer */
	Payload     string `json:"payload"`
}

func Serve() {
	fmt.Println("Starting server...")
	fmt.Println("-------------------------------")
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleConnection)
	srv := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleConnection(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error! Cannot read request body")
		w.Write([]byte("Error! Cannot read request body."))
		return
	}

	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		fmt.Println("Error! Cannot parse request body.")
		fmt.Println(err.Error())
		w.Write([]byte("Error! Cannot parse request body."))
	}

	/* If exit node ... */
	if message.ExitAddress != nil {
		w.Write([]byte(*message.ExitAddress + "\n"))
		fmt.Println("Exit node!")
		return
	}

	/* If relay node ... */
	if message.Next != nil {
		w.Write([]byte(*message.Next + "\n"))
		fmt.Println("Relay node!")
		return
	}
}
