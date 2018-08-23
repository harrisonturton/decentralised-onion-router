package server

import (
	"onion-router/comm"
	"onion-router/exit"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	var message comm.Message
	if err := json.Unmarshal(body, &message); err != nil {
		fmt.Println("Error! Cannot parse request body.")
		fmt.Println(err.Error())
		w.Write([]byte("Error! Cannot parse request body."))
	}

	/* If exit node ... */
	if message.ExitAddress != nil {
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
		fmt.Println("Exit node!")
		return
	}

	/* If relay node ... */
	if message.Next != nil {
		relayMessage := comm.RelayMessage{
			Next: *message.Next,
			Payload: message.Payload,
		}
		w.Write([]byte(relayMessage.Next + "\n"))
		fmt.Println("Relay node!")
		return
	}
}
