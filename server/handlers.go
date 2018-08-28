package server

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"onion-router/peers"
)

func makeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(peers.Routes[peers.Announce], announceHandler)
	mux.HandleFunc(peers.Routes[peers.Route], routeHandler)
	mux.HandleFunc(peers.Routes[peers.Exit], exitHandler)
	mux.HandleFunc(peers.Routes[peers.Relay], relayHandler)
	return mux
}

func announceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if len(peers.Peers) == 0 {
		return
	}
	var addrs []peers.NodeAddr
	for _, peer := range peers.Peers {
		addrs = append(addrs, peers.NodeAddr{
			Address: peer.Address,
			Port:    peer.Port,
		})
	}
	resp, err := json.Marshal(addrs)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to marshal peers in announce response"))
		return
	}
	w.Write(resp)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "Hello route!")
}

func exitHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "Hello exit!")
}

func relayHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "Hello relay!")
}
